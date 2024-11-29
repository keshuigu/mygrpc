# 多路复用

一个 `grpc.ClientConn` 可以被两个存根共享，两个服务可以共享一个 `grpc.Server`。此示例演示了如何执行这两种类型的共享。

```
go run server/main.go
```

```
go run client/main.go
```
