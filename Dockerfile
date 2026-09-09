FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -o next-erp-api ./cmd/api

FROM alpine:3.22

RUN apk upgrade --no-cache && \
    addgroup -S appgroup && \
    adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/next-erp-api .

RUN chown appuser:appgroup /app/next-erp-api

USER appuser

EXPOSE 8081

CMD ["./next-erp-api"]