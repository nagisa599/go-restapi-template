# スキーマ変更手順

### OpenAPI のスキーマを変更する

- [app/schema](../app/schema)の中身を変更する

- `Swagger Viewer`の VsCode を拡張機能を入れることをおすすめする

### OpenAPI.yaml ファイルを生成する

```bash
make gen-openapi
```

- [app/schema](../app/schema)は、分割ファイルとなっているので 1 つの OpenAPI に統合する

### OpenAPI ファイルから yaml ファイルを生成する

```bash
make gen-go
```
