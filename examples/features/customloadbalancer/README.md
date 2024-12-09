# TODO
# 自定义负载均衡器

此示例展示了如何在 `ClientConn` 中部署自定义负载均衡器。

## 试一试

```
go run server/main.go
```

```
go run client/main.go
```

## 解释

两个回显服务器在 "localhost:20000" 和 "localhost:20001" 上提供服务。它们将在响应中包含其服务地址。因此，"localhost:20001" 上的服务器将回复 RPC 并显示 `this is examples/customloadbalancing (from localhost:20001)`。

创建了一个客户端，连接到这两个服务器（它们从名称解析器中获取两个独立端点的服务器地址）。客户端配置了在服务配置中指定的负载均衡器，在此示例中是`custom_round_robin`

### custom_round_robin

客户端配置为使用`custom_round_robin`。`custom_round_robin` 为它接收到的每个端点创建一个 `pick first` 子负载均衡器。它等待两个 `pick first` 子负载均衡器都准备就绪，然后依赖第一个 `pick first` 子负载均衡器的选择器，选择连接到 `localhost:20000`，除了每隔 `chooseSecond` 次，它依赖第二个 `pick first` 子负载均衡器的选择器，选择连接到 `localhost:20001`（或反之亦然）。

`custom_round_robin` 被编写为一个委托策略，包装了每个接收到的端点的 `pick_first `负载均衡器。这是用户编写自定义负载均衡器的预期方式，因为 `pick first` 将包含许多有用的功能，例如粘性瞬态故障、`Happy Eyeballs` 和健康检查。


```
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50051)
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50051)
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50050)
this is examples/customloadbalancing (from localhost:50051)
```
