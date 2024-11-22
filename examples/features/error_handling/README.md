# Description

此示例演示了在gRPC中基本的RPC错误处理。

# Run the sample code

运行服务器，如果RPC请求的Name字段为空，服务器将返回错误。

```sh
$ go run ./server/main.go
```

然后在另一个终端中运行客户端，客户端将进行两次请求：一次请求的Name字段为空，另一次请求的Name字段填充了由os/user提供的当前用户名。


```sh
$ go run ./client/main.go
```

它应该打印从服务器接收到的状态码。
