FROM golang:1.24.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /main .
COPY ./config/config.yaml ./config/
COPY ./config/.env ./config/
COPY ./db/migrations/ /app/db/migrations/
RUN chmod -R 755 /app/db/migrations

RUN ls -la /app/db/migrations/

#ENV CONFIG_PATH=/app/config/config.yaml
ENV CONFIG_PATH=/app/config/.env
ENV MIGRATIONS_PATH="file:///app/db/migrations"

EXPOSE 8081

CMD ["./main"]