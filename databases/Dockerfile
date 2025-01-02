FROM golang:1.18

# 作業ディレクトリを設定
WORKDIR /app

# Gooseをインストール
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# マイグレーションファイルをコピー
COPY ./migrations /app/migrations

# デフォルトコマンドを設定（例：PostgreSQLへの接続情報を指定）
CMD ["goose", "postgresql", "dbname=postgres user=postgres password=secret host=db port=5432 sslmode=disable", "up"]