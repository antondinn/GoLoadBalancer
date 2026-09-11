FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend ./cmd/backend

FROM alpine

COPY --from=builder /backend /backend

EXPOSE 8080

ENTRYPOINT ["/backend"]