FROM --platform=linux/arm64 golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -o api ./cmd/api

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 \
    go build -o worker ./cmd/worker

EXPOSE 8080

CMD ["./api"]