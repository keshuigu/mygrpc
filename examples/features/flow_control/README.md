# 流控

流控是 gRPC 中的一项功能，它防止发送方在流上写入超过接收方能够处理的数据量。此功能在客户端和服务器端的行为相同。由于 gRPC-Go 使用阻塞式 API 进行流操作，当流的流控限制达到时，流控反馈通过简单地阻塞流上的发送操作来实现。当接收方从流中读取了足够的数据后，发送操作将自动解除阻塞。流控会根据连接的带宽延迟积（BDP）自动配置，以确保缓冲区的大小是允许流在接收方以最大速度读取时实现最大吞吐量所需的最小大小。

## Try it

```
go run ./server
```

```
go run ./client
```

## Example explanation

示例客户端和服务器的编写目的是通过在另一方未接收时故意发送消息来演示阻塞。示例中的双向回显流开始时，客户端会发送消息，直到检测到它已被阻塞（利用另一个 goroutine）。服务器会休眠 2 秒以允许这种情况发生。然后服务器将读取所有这些消息，并交换客户端和服务器的角色，使服务器尝试连续发送，而客户端休眠。客户端休眠 2 秒后，它将再次读取以解除服务器的阻塞。服务器将检测到它已被阻塞，并在解除阻塞后结束流。

### 预期输出

client
```
2023/09/19 15:49:49 New stream began.
2023/09/19 15:49:50 Sending is blocked.
2023/09/19 15:49:51 Sent 25 messages.
2023/09/19 15:49:53 Read 25 messages.
2023/09/19 15:49:53 Stream ended successfully.
```

server

```
2023/09/19 15:49:49 New stream began.
2023/09/19 15:49:51 Read 25 messages.
2023/09/19 15:49:52 Sending is blocked.
2023/09/19 15:49:53 Sent 25 messages.
2023/09/19 15:49:53 Stream ended successfully.
```