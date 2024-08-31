# Usando a imagem oficial do Go como base
FROM golang:1.20

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do projeto para o diretório de trabalho
COPY . .

# Instalar SQLite3
RUN apt-get update && apt-get install -y sqlite3 libsqlite3-dev

# Compilar o código Go para server e client
WORKDIR /app/Server
RUN go mod tidy
RUN go build -o /app/server

WORKDIR /app/Client
RUN go mod tidy
RUN go build -o /app/client

# Expor a porta do server.go
EXPOSE 8080

# Comando para iniciar o server.go
CMD ["/app/server"]
