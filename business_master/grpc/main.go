package main

import (
	lightweight "business_master/pb"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

// server继承自动生成的服务类
type server struct {
	lightweight.UnimplementedGymServer
}

// 实现接口
func (s *server) BodyBuilding(ctx context.Context, in *lightweight.Person) (*lightweight.Reply, error) {
	fmt.Println("正在健身 ， 动作 ", in.Name, in.Actions)
	return &lightweight.Reply{Code: 0, Msg: "ok"}, nil
}
func main() {
	// 监听端口
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("failed to listen :", err)
		return
	}
	// 创建一个grpc服务
	s := grpc.NewServer()
	// 注册服务
	lightweight.RegisterGymServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Error("failed to serve: ", err)
	}
}
