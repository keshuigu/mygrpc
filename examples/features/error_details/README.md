# Description

此示例演示了在gRPC错误中使用状态详情。

# Run the sample code

运行服务器:

```sh
$ go run server/main.go
```
然后在另一个终端中运行客户端：

```sh
$ go run client/main.go
```

它应该成功并打印从服务器接收到的问候语。然后再次运行客户端：

```sh
$ go run client/main.go
```

这次，它应该失败，并打印从服务器接收到的错误状态详情