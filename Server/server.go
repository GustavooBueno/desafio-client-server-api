package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	http.HandleFunc("/cotacao", handleCotacao)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleCotacao(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Timeout para obter cotação
	cotacaoCtx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	cotacao, err := obterCotacao(cotacaoCtx)
	if err != nil {
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		log.Println("Erro ao obter cotação:", err)
		return
	}

	// Timeout para salvar no banco de dados
	dbCtx, cancelDb := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancelDb()

	err = salvarCotacao(dbCtx, cotacao.Bid)
	if err != nil {
		http.Error(w, "Erro ao salvar cotação", http.StatusInternalServerError)
		log.Println("Erro ao salvar cotação:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"bid": cotacao.Bid})
}

func obterCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	cotacao := result["USDBRL"]
	return &cotacao, nil
}

func salvarCotacao(ctx context.Context, bid string) error {
	db, err := sql.Open("sqlite3", "./cotacoes.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS cotacoes (id INTEGER PRIMARY KEY, bid TEXT, data TIMESTAMP)")
	if err != nil {
		return err
	}

	_, err = db.ExecContext(ctx, "INSERT INTO cotacoes (bid, data) VALUES (?, ?)", bid, time.Now())
	if err != nil {
		return err
	}

	return nil
}
