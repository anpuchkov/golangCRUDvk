FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy
RUN go build -o app ./cmd/main.go


FROM postgres:latest

COPY sql/database.sql /docker-entrypoint-initdb.d/

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD 1234
ENV POSTGRES_DB films


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

COPY configs/config.yaml ./configs/config.yaml

COPY model/db/ ./model/db/

EXPOSE 8080

CMD ["./app"]
