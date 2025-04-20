# ビルドステージ
FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache tzdata

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o bin ./main.go

# 実行ステージ
FROM gcr.io/distroless/static:nonroot

COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
ENV TZ=Asia/Tokyo

USER nonroot:nonroot

WORKDIR /app

COPY --from=builder /app/bin /app/

EXPOSE 8080

CMD ["/app/bin"]
