# Client-Server-API - Cotação do dólar

## Requisitos
1. Entrega de 2 sistemas:
    - client.go
    - server.go

2. O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.

3. O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL 
e em seguida deverá retornar no formato JSON o resultado para o cliente.

4. Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, 
sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms 
e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.

5. O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). 
Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.

6. Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.

7. O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}

8. O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.

## Como rodar
- Certifique de ter Go instalado na máquina
- Clone o repositório: git clone https://github.com/GustavooBueno/desafio-client-server-api.git
- cd desafio-client-server-api
- Rodar servidor:
    - cd Server
    - go run server.go
- Rodar cliente:
    - cd Client
    - go run client.go
- No diretório Client será criado um arquivo cotacao.txt
- Para verificar o Banco de Dados:
    - sqlite3 Server/cotacoes.db
    - SELECT * FROM cotacoes;
