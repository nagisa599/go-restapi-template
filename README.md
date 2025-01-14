# go-restapi-template

## 初期設定

```bash
# swagger-cliのインストール
npm install -g swagger-cli
```

## テスト設計

初期設定

```bash
# gomockをインストール
make install
```

テストは以下の 3 つで行う

### repository のテスト

`go-sqlmock`を利用
手順

1. 〇〇\_repository.go で copilot を使って`@workspace/test`を利用してテストを作成する
2. 微修正は、[以下のファイル](./app/internal/domain/repository/user_repository_test.go)を参考に

### usecase のテスト

`gomock`を利用

1. repository の mock を作成する.[makefile を参考に](./makefile)
2. 〇〇\_usecase.go で copilot を使って`@workspace/test`を利用してテストを作成する

### handler のテスト

`gomock`を利用

1. usecase の mock を作成する.[makefile を参考に](./makefile)
2. 〇〇\_handler.go で copilot を使って`@workspace/test`を利用してテストを作成する

## e2e テスト

## エラー設計

| Error Type            | Code | Message               |
| --------------------- | ---- | --------------------- |
| Bad Request           | 400  | Bad request           |
| Unauthorized          | 401  | Unauthorized          |
| Forbidden             | 403  | Forbidden             |
| Not Found             | 404  | Not found             |
| Internal Server Error | 500  | Internal server error |

## 参考記事

https://qiita.com/unsoluble_sugar/items/b080a16701946fcfce70

## エラーハンドリング

- usecase でエラーを投げられない
