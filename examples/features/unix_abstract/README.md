# Unix 抽象套接字

此示例展示了如何启动一个监听 Unix 抽象套接字的 gRPC 服务器以及如何让 gRPC 客户端连接到它。

## 什么是 Unix 抽象套接字

抽象套接字地址与常规 Unix 套接字的区别在于地址的第一个字节是空字节（'\0'）。该地址与文件系统路径名没有关联。

## 试一试

```
go run server/main.go
```

```
go run client/main.go
```

## 解释

此示例中的 gRPC 服务器监听一个以空字节开头且网络为 `unix` 的地址。客户端使用 `unix-abstract` 方案，端点设置为不带空字节的抽象 Unix 套接字地址。`unix` 解析器负责在客户端添加空字节。更多详情请参见 https://github.com/grpc/grpc/blob/master/doc/naming.md。