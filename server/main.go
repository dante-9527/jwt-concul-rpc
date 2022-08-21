package main

// jwt-consul-rpc-server

import (
	"context"
	"jwt-consul-rpc-server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

func (s *server) QueryInfo(ctx context.Context, in *pb.QueryRequest) (*pb.QueryResponse, error) {
	reply := new(pb.QueryResponse)
	// 根据传递进来的id或者Name去数据库进行查询
	// 获取metadata，校验jwt

	// 获取数据进行校验 id

	// 连接mysql进行查询

	// 返回结果
	reply.Code = 0
	reply.SAge = 18
	reply.SName = "dante"
	return reply, nil

}

func main() {
	// 监听本地的8972端口
	lis, err := net.Listen("tcp", ":8972")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()                  // 创建gRPC服务器
	pb.RegisterGetInfoServer(s, &server{}) // 在gRPC服务端注册服务
	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}

}
