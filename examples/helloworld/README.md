# gRPC Hello World

1. 生成*pb.go文件 安装protobuf-compiler

  ```console
  $ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    helloworld/helloworld.proto
  ```

2. 启动服务器

  ```console
  $ go run greeter_server/main.go &
  ```

3. 启动客户端

  ```console
  $ go run greeter_client/main.go
  Greeting: Hello world
  ```


> From gRpc
> For more details (including instructions for making a small change to the
> example code) or if you're having trouble running this example, see QuickStart: https://grpc.io/docs/languages/go/quickstart