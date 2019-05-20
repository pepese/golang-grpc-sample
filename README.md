```bash
$ GO111MODULE=on go mod init
$ mkdir proto
$ mkdir proto/dest
$ mkdir proto/dest/helloworld
$ touch proto/helloworld.proto
$ protoc -I=./proto --go_out=plugins=grpc:./proto/dest proto/helloworld.proto
```