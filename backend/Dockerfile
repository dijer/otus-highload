FROM golang:1.23 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app ./cmd/app


FROM golang:1.23
WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/config.toml ./config.toml

ENV CONFIG_PATH=/app/config.toml
ENV MIGRATIONS_PATH=/app/migrations

CMD ["./app"]
