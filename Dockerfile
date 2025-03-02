FROM golang:1.22.1-alpine AS builder

ENV GOOS linux

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/app ./cmd/.

FROM alpine:latest AS go-app

WORKDIR /app

COPY --from=builder /app/app /app/app

COPY --from=builder /app/var /app/var

ENTRYPOINT ["/app/app"]