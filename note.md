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
10. `runtime.Caller()`是Go语言标准库中的一个函数，用于获取当前调用栈的信息。它可以返回调用者的文件名、行号和函数信息。这个函数在调试和日志记录中非常有用，可以帮助开发者了解代码的执行路径
    1. 参数skip：表示需要获取的栈帧数。0表示当前函数(Caller的调用者)，1表示当前函数的调用者，以此类推
    2. 返回值
       1. pc：程序计数器，表示当前执行的指令地址。
       2. file：调用者所在的文件名。
       3. line：调用者所在的行号。
       4. ok：如果信息获取成功，返回true，否则返回false。
11. `func init()`是特殊函数，在调用包时执行
    1.  自动调用：init函数在包初始化时自动调用，无需显式调用。
    2.  无参数和返回值：init函数不能有参数和返回值。
    3.  每个包可以有多个init函数：一个包中可以有多个init函数，甚至在同一个文件中也可以有多个init函数。它们的执行顺序是按照它们在文件中出现的顺序。
    4.  每个文件可以有多个init函数：一个文件中可以有多个init函数，执行顺序是按照它们在文件中出现的顺序。
12. 基础是http2协议，需要学习
13. keepalive中的 stream是 proto的,指的是没有活动的RPC流也发送ping
14. go中反引号代表原始字符串，即不需要转义
15. metadata client发送metadata，而server区分头部和尾部
16. 一元rpc可以用calloption的方法注册接受header和trailer的metadata，流式的不可以，需要用stream.Header()接受
17. make(chan struct{}) 是 Go 语言中的一种语法，用于创建一个空的（零大小）通道。这个通道的元素类型是 `struct{}`，表示一个空的结构体。使用 `struct{}` 作为通道的元素类型有几个优点：**节省内存**：`struct{}` 是一个零大小的类型，因此使用它作为通道的元素类型不会占用额外的内存。**表示信号**：使用 `struct{}` 通道通常用于信号传递，而不是传递实际的数据。它可以用来表示事件的发生或完成。
18. TIME_WAIT TCP中接收到FIN发送ACK后等待关闭，防止对方未接收到ACK而需要重传FIN
19. defer后进先出
20. ```go
    func main() {
        var i interface{} = "hello"
        var s interface{}
        // var ok bool
        if s, ok := i.(string); ok {
            println(s)
        } else {
            println("no")
        }
        println(s) // 不受if的赋值影响
    }
    ```