# 重试

此示例展示了如何在 gRPC 客户端上启用和配置重试。

## 文档

[客户端重试支持的 gRFC](https://github.com/grpc/proposal/blob/master/A6-client-retries.md)

## 试一试

此示例包括一个服务实现，该实现会使请求三次失败并返回状态码 `Unavailable`，然后通过第四次请求。客户端配置为在收到 `Unavailable` 状态码时进行四次重试尝试。

首先启动服务器：

```bash
go run server/main.go
```

然后运行客户端：

```bash
go run client/main.go
```

## 用法

### 定义您的重试策略

通过服务配置启用重试，可以由名称解析器或 DialOption 提供（如下所述）。在下面的配置中，我们为 "grpc.example.echo.Echo" 方法设置了重试策略。

MaxAttempts：在失败之前尝试 RPC 的次数。
InitialBackoff, MaxBackoff, BackoffMultiplier：配置尝试之间的延迟。
RetryableStatusCodes：仅在收到这些状态码时重试。


```go
        var retryPolicy = `{
            "methodConfig": [{
                // config per method or all methods under service
                "name": [{"service": "grpc.examples.echo.Echo"}],

                "retryPolicy": {
                    "MaxAttempts": 4,
                    "InitialBackoff": ".01s",
                    "MaxBackoff": ".01s",
                    "BackoffMultiplier": 1.0,
                    // this value is grpc code
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`
```

### 作为 DialOption 提供重试策略

要使用上述服务配置，请将其与 grpc.WithDefaultServiceConfig 一起传递给 grpc.NewClient。


```go
conn, err := grpc.NewClient(ctx,grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
```