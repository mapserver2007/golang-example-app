# golang-example-app

## サーバ起動
```sh
$> rake
$> http://localhost:3000/v1/users -i # openapi-server
$> http://localhost:3001/v1/users -i # grpc-server
```

## protoの拡張
go getコマンドでprotoファイルの拡張を使う場合にFile not foundエラーになるときは
`/usr/local/Cellar/protobuf/${version}/include/validate/`配下にファイル配置する。

例：/usr/local/Cellar/protobuf/3.12.3/include/google/api/annotations.proto