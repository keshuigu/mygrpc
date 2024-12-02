package statshandler

import (
	"context"
	"log"
	"net"
	"path/filepath"

	"google.golang.org/grpc/stats"
)

// 实现 [stats.Handler](https://pkg.go.dev/google.golang.org/grpc/stats#Handler)
type Handler struct{}

type connStatCtxKey struct{}

// TagConn 可以将一些信息附加到给定的上下文中。
// 用于此连接的 HandleConn 中的上下文将从返回的上下文派生。
// 在 gRPC 客户端中：
// 用于此连接上的 RPC 的 HandleRPC 中的上下文将是用户的上下文，而不是从这里返回的上下文派生的。
// 在 gRPC 服务器中：
// 用于此连接上的 RPC 的 HandleRPC 中的上下文将从这里返回的上下文派生。
func (st *Handler) TagConn(ctx context.Context, stat *stats.ConnTagInfo) context.Context {
	log.Printf("[TagConn] [%T]: %+[1]v", stat)
	// log.Printf("[TagConn] [%T]: %+v", stat, stat) // same
	//%T 打印stat类型
	//%+[1]v "+" -> 打印字段名 "[1]" -> 使用第一个参数
	return context.WithValue(ctx, connStatCtxKey{}, stat)
}

func (st *Handler) HandleConn(ctx context.Context, stat stats.ConnStats) {
	var rAddr net.Addr
	// 接口类型断言
	// any.(type)
	// ==> s -> type
	// ==> ok -> bool
	if s, ok := ctx.Value(connStatCtxKey{}).(*stats.ConnTagInfo); ok {
		rAddr = s.RemoteAddr
	}
	if stat.IsClient() {
		log.Printf("[server addr: %s] [HandleConn] [%T]: %+[2]v", rAddr, stat)
	} else {
		log.Printf("[client addr: %s] [HandleConn] [%T]: %+[2]v", rAddr, stat)
	}
}

type rpcStatCtxKey struct{}

func (st *Handler) TagRPC(ctx context.Context, stat *stats.RPCTagInfo) context.Context {
	log.Printf("[TagRPC] [%T]: %+[1]v", stat)
	return context.WithValue(ctx, rpcStatCtxKey{}, stat)
}

func (st *Handler) HandleRPC(ctx context.Context, stat stats.RPCStats) {
	var sMethod string
	if s, ok := ctx.Value(rpcStatCtxKey{}).(*stats.RPCTagInfo); ok {
		sMethod = filepath.Base(s.FullMethodName)
	}

	var cAddr net.Addr
	if s, ok := ctx.Value(connStatCtxKey{}).(*stats.ConnTagInfo); ok {
		cAddr = s.RemoteAddr
	}

	if stat.IsClient() {
		log.Printf("[server method: %s] [HandleRPC] [%T]: %+[2]v", sMethod, stat)
	} else {
		log.Printf("[client addr: %s] [HandleRPC] [%T]: %+[2]v", cAddr, stat)
	}

}

// New returns a new implementation of [stats.Handler](https://pkg.go.dev/google.golang.org/grpc/stats#Handler) interface.
func New() *Handler {
	return &Handler{}
}
