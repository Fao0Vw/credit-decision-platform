# Stage 1 — build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# сначала зависимости (кеширование)
COPY go.mod  ./
RUN go mod download

# потом код
COPY . .

# сборка бинарника
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go

# Stage 2 — runtime
FROM alpine:3.19

WORKDIR /app

# копируем только бинарник
COPY --from=builder /app/app .

# порт (для документации)
EXPOSE 8080

# запуск
CMD ["./app"]