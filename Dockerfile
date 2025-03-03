FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o app ./cmd/.

FROM alpine:latest AS go-app

ENV NUTRIAI_PORT=8080

ENV API_KEY=YOR_KEY

ENV REDIS_HOST=redis

ENV REDIS_PORT=6379

WORKDIR /app

COPY --from=builder /app/app /app/app

COPY --from=builder /app/var /app/var

CMD ["./app"]