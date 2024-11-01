# gRPC Hello World

1. 启动服务器

  ```console
  $ $(go env GOPATH)/bin/greeter_server &
  ```
2. 启动客户端

  ```console
  $ $(go env GOPATH)/bin/greeter_client
  Greeting: Hello world
  ```


> From gRpc
> For more details (including instructions for making a small change to the
> example code) or if you're having trouble running this example, see QuickStart: https://grpc.io/docs/languages/go/quickstart