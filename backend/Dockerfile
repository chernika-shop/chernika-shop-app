FROM golang:1.25-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/air-verse/air@latest

# Устанавливаем Atlas через бинарник (из-за проблем с компиляцией в Go 1.25)
RUN apk add --no-cache curl bash && \
    curl -sSf https://atlasgo.sh | bash -s -- --dir /usr/local/bin && \
    chmod +x /usr/local/bin/atlas && \
    apk del curl bash

ENV PATH=$PATH:/go/bin

COPY . .

RUN swag init -g cmd/server/main.go -o api/docs

EXPOSE 8000

CMD ["air", "-c", ".air.toml"]