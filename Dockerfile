FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o NutriAI ./cmd/.

FROM alpine:latest AS NutriAI

ENV NUTRIAI_PORT=8080

ENV API_KEY=ZDMxOTdmNjUtMmY3MS00MTdjLThkY2YtODljY2RiZGI1ZDZkOjVlMmM3OWYxLTUwNDQtNDRkNi05NTY1LTA3NzBlNTkyMWNmMQ==

ENV REDIS_HOST=redis

ENV REDIS_PORT=6379

ENV LOG_FILE=./var/log/app.log

WORKDIR /app

COPY --from=builder /app/NutriAI /app/NutriAI

COPY --from=builder /app/var /app/var

CMD ["./NutriAI"]