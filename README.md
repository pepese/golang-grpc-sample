```bash
$ pwd
$GOPATH/src/github.com/pepese/golang-grpc-sample
$ GO111MODULE=on
$ go mod init
$ mkdir proto
$ mkdir proto/dest
$ mkdir proto/dest/helloworld
$ touch proto/helloworld.proto # $GOPATH/src/google.golang.org/grpc/examples のサンプルを持ってくる
$ protoc -I=./proto --go_out=plugins=grpc:./proto/dest proto/*.proto
```

`--go_out` オプションで指定された出力先に対して、 .proto ファイルの `go_package` で指定した相対パスのディレクトリ構成で出力される。

サーバサイドの実行。

```bash
$ go run main.go
```

クライアントの実行

```bash
$ go run client.go
```