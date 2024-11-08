1. `*.proto`是用于定义 Protocol Buffers 数据结构和服务的文件格式。Protocol Buffers 是一种语言无关、平台无关、可扩展的序列化机制
2. 如果您想在 RPC（远程过程调用）系统中使用消息类型，您可以在 `.proto` 文件中定义 RPC 服务接口，`protocol buffer` 编译器将以您选择的语言生成服务接口代码和存根。因此，例如，如果您想使用接受 SearchRequest 并返回 SearchResponse 的方法定义 RPC 服务，则可以在 .proto 文件中定义它，如下所示：
   ```proto
   service SearchService {
    rpc Search(SearchRequest) returns (SearchResponse);
   }
   ```
3. flag包用于解析命令行参数
4. var 在函数外声明全局变量
5. if语句中声明的变量，作用范围只在if块中,包括elseif和else分支
6. map需要用make初始化`routeNotes := make(map[string][]*pb.RouteNote)`
7. make 初始化切片，映射，通道
8.	`grpcServer := grpc.NewServer(opts...)`，其中`...`将切片解包为可变参数列表，传递给函数。
9.	`ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)`10s超时上下文，10s后发送ctx.Done()到chan
