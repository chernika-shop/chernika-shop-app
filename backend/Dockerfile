FROM golang:1.25-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go install github.com/air-verse/air@latest

ENV PATH=$PATH:/go/bin

COPY . .

RUN swag init -g cmd/server/main.go -o api/docs

# Открыть порт
EXPOSE 8000

CMD ["air", "-c", ".air.toml"]