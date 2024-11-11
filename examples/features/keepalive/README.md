> Keepalive
>此示例说明了如何设置客户端的保活ping和服务器端的保活ping强制执行以及连接空闲设置。有关这些设置的更多详细信息，请参见[英文完整文档](https://github.com/grpc/grpc-go/tree/master/Documentation/keepalive.md).
> ```
> go run server/main.go
> GODEBUG=http2debug=2 go run client/main.go
> ```


# Keepalive

gRPC在传输层发送HTTP2 ping以检测连接是否断开。如果在一定时间内未收到对方的ping响应，连接将被关闭。请注意，只有在连接上没有活动时才需要ping。

有关如何配置Keepalive，请参见https://godoc.org/google.golang.org/grpc/keepalive中的选项。

## 为什么我需要这个？

Keepalive可以用于检测TCP级别的连接故障。一个特定的情况是，当TCP连接丢包（包括FIN包）时，系统TCP超时（可能长达30分钟）才能检测到这个故障。Keepalive可以让gRPC更快地检测到这个故障。

另一个用途是（顾名思义）保持连接活跃。例如，在L4代理配置为杀死“空闲”连接的情况下，发送ping可以使连接不“空闲”。

## 我应该设置什么？

对于大多数用户来说，将[客户端参数](https://godoc.org/google.golang.org/grpc/keepalive)设置为[拨号选项](https://godoc.org/google.golang.org/grpc#WithKeepaliveParams)应该就足够了。

## 会发生什么？

（这里描述的行为特定于gRPC-go，在其他语言中可能略有不同。）

当连接上没有活动时（请注意，当没有消息发送时，正在进行的流会导致__no activity__），在`Time`之后，客户端将发送一个ping，服务器在收到ping时将发送一个ping响应。客户端将等待`Timeout`，并检查在此期间连接上是否有任何活动（ping响应是一种活动）。

## 服务器端呢？

服务器端有与客户端类似的`Time`和`Timeout`设置。服务器还可以配置连接的最大年龄。有关详细信息，请参见[服务器参数](https://godoc.org/google.golang.org/grpc/keepalive#ServerParameters)

### 强制策略

[强制策略](https://godoc.org/google.golang.org/grpc/keepalive#EnforcementPolicy) 服务器端的一种特殊设置，用于保护服务器免受恶意或行为不当的客户端的影响。

当检测到不良行为时，服务器将发送带有`ENHANCE_YOUR_CALM`的`GOAWAY`并关闭连接：

- 客户端发送ping过于频繁
- 客户端在没有流的情况下发送ping，而这是服务器配置不允许的
