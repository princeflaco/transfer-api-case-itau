FROM golang:1.22.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o transfer-api-case-itau ./main.go

# Possivéis variáveis de ambiente opcionais
#ENV PORT=8080
#ENV TIMEOUT=30
#ENV APP_NAME="Transfer API"

EXPOSE 8080

CMD ["./transfer-api-case-itau"]
