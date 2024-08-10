# Dockerfile
# ビルドステージ
FROM golang:1.16 AS builder

WORKDIR /app

# go.mod と go.sum ファイルをコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# 依存関係を整理
RUN go mod tidy

# main.go を含むディレクトリに移動
WORKDIR /app/cmd/server

# アプリケーションをビルド
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# 実行ステージ
FROM alpine:latest

WORKDIR /app

# ビルドステージからバイナリをコピー
COPY --from=builder /app/main .

# 環境変数 PORT を設定
ENV PORT=8080

# ポートを公開
EXPOSE 8080

# アプリケーションを起動
CMD ["/app/main"]