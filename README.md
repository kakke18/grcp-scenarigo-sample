# grpc-scenarigo-sample
[scenarigo](https://github.com/zoncoen/scenarigo)でgRPCのE2Eテストを実装

## 実行方法
### サーバ起動
```shell
# grpc-scenarigo-sample配下
make run
```

### E2Eテスト実行
```shell
# サーバ起動中に別Terminalで実行する想定
# grpc-scenarigo-sample配下
cd e2e
make init # 必要なコマンドのインストール
make build # pluginのビルド
make test
```

## ディレクトリ構成
### cmd
- server/main.go
  - gPRCサーバの起動処理を実装したファイル

### e2e
- gen
  - プラグインのビルドファイル（`.so`）を格納するディレクトリ
- plugins
  - 自作したプラグインを格納するディレクトリ
  - grpc
    - scenarigoから使えるように、gRPC Clientの生成関数を実装
    - pb
      - gPRC Clientの生成関数のために、`.proto`ファイルから生成される`.go`ファイルを格納する
- scenarios
  - シナリオテストを記述したyamlファイルを格納するディレクトリ
- Makefile
  - 便利なコマンドをまとめたファイル
- scenarigo.yaml
  - scenarigoの設定ファイル

### internal
- dao
  - データベースへのアクセスを実装するパッケージ
  - 今回はDBを用意するのが面倒だったので、インメモリで実装している
- model
  - アプリケーションで扱うデータの構造体を定義するパッケージ
- service
  - アプリケーションのビジネスロジックを実装するパッケージ

### pb
- `.proto`ファイルから生成される`.go`ファイルを格納するディレクトリ

### proto
- echo/v1/echo.proto
  - EchoServiceの定義ファイル
- user/v1/user.proto
  - UserServiceの定義ファイル
- buf.gen.yaml
  - [buf](https://github.com/bufbuild/buf)の設定ファイル。詳細は省略
- buf.yaml
  - [buf](https://github.com/bufbuild/buf)の設定ファイル。詳細は省略
- Makefile
    - 便利なコマンドをまとめたファイル
