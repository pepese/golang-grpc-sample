```bash
$ pwd
$GOPATH/src/github.com/pepese/golang-grpc-sample
$ GO111MODULE=on
$ go mod init
$ mkdir proto
$ mkdir proto/dest
$ mkdir proto/dest/helloworld
$ touch proto/helloworld.proto # $GOPATH/src/google.golang.org/grpc/examples のサンプルを持ってくる
$ protoc -I=./proto --go_out=plugins=grpc:./proto/dest/helloworld proto/*.proto
```

サーバサイドの実行。

```bash
$ go run main.go
```

クライアントの実行

```bash
$ go run client.go
```