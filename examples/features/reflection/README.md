# 反射

此示例展示了如何在gRPC服务器上注册反射。

参考：https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md

# 试一试

```go
go run server/main.go
```

有多个现有的反射客户端。

`grpcurl`: https://github.com/fullstorydev/grpcurl.

```zsh
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 list helloworld.Greeter
grpcurl -plaintext localhost:50051 describe helloworld.Greeter
grpcurl -plaintext localhost:50051 describe mygrpc.examples.echo.Echo
grpcurl -plaintext localhost:50051 describe helloworld.HelloRequest
grpcurl -plaintext -format text -d 'name: "gRPCurl"' localhost:50051 helloworld.Greeter.SayHello
```