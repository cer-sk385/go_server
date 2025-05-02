# ビルドステージ
FROM golang:1.24-alpine AS builder

WORKDIR /app

# 依存関係のファイルをコピー
COPY go.mod ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 実行ステージ
FROM alpine:latest

WORKDIR /app

# ビルドステージから実行ファイルをコピー
COPY --from=builder /app/main .

# ポート8080を公開
EXPOSE 8080

# アプリケーションを実行
CMD ["./main"]