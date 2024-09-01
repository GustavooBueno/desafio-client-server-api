package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// Timeout 300ms para obter cotação
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		log.Fatal("Erro ao criar requisição:", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Erro ao obter cotação:", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal("Erro ao decodificar resposta:", err)
	}

	bid, ok := result["bid"]
	if !ok {
		log.Fatal("Erro: campo 'bid' não encontrado na resposta")
	}

	fmt.Printf("Cotação atual do dólar: %s\n", bid)

	// Salvar cotação em arquivo
	err = salvarCotacaoEmArquivo(bid)
	if err != nil {
		log.Fatal("Erro ao salvar cotação em arquivo:", err)
	}
}

func salvarCotacaoEmArquivo(bid string) error {
	content := fmt.Sprintf("Dólar: %s", bid)
	return ioutil.WriteFile("cotacao.txt", []byte(content), 0644)
}
