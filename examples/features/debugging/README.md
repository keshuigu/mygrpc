# 调试
目前，gRPC提供了两种主要的工具来帮助用户调试问题，分别是logging和Channelz。

## logging

gRPC在关键路径上放置了大量的日志记录工具，以帮助用户调试问题。日志级别文档描述了每个日志级别在gRPC上下文中的含义。

要打开调试日志，请使用以下环境变量运行代码： GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info。

## Channelz

我们还提供了一个运行时调试工具Channelz，以帮助用户进行实时调试。

有关如何使用Channelz服务调试实时程序的详细信息，请参阅[Channelz博客文章](https://grpc.io/blog/a-short-introduction-to-channelz/)。

需要[前端程序](https://github.com/grpc/grpc-experiments/tree/master)

## 试一试

该示例展示了logging和Channelz如何帮助调试。有关完整解释，请参阅上面链接的Channelz博客文章。

