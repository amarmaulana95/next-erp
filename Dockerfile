FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o next-erp-api ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/next-erp-api .

EXPOSE 8081

CMD ["./next-erp-api"]
