# ビルドステージ
FROM golang:1.25 AS builder
WORKDIR /app

# go.mod/go.sumを先にコピーして依存パッケージをキャッシュする
COPY go.mod go.sum ./
RUN go mod download

# ソースコード全体をコピー
COPY . .

# CGO無効で純粋なGoバイナリを生成（distrolessで動かすため）
RUN CGO_ENABLED=0 go build -o server ./main.go

# 実行ステージ（TLS証明書入りの軽量イメージ）
FROM gcr.io/distroless/static
COPY --from=builder /app/server /server
CMD ["/server"]
