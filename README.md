# golang-example-app

![build](https://github.com/mapserver2007/golang-example-app/workflows/build/badge.svg)

## サーバ起動
```sh
$> rake
$> http://localhost:4001/v1/users_and_items -i # grpc-server
```

## 試したこと
* サーバの起動
  * grpc-gatewayサーバ
  * grpcサーバ(本体)
  * grpcサーバ(サービス1)
  * grpcサーバ(サービス2)
* サーバ間gRPC通信
* DB処理
  * gorp
* ログ出力
  * logrus
* gatewayサーバ
  * バリデーション
    * RESTAPIのバリデーション(protobufの拡張)
  * エラーハンドリング
    * RESTAPIのエラーレスポンスをカスタマイズ
* OpenAPI

## protoの拡張
go getコマンドでprotoファイルの拡張を使う場合にFile not foundエラーになるときは
`/usr/local/Cellar/protobuf/${version}/include/validate/`配下にファイル配置する。

例：/usr/local/Cellar/protobuf/3.12.3/include/google/api/annotations.proto