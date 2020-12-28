package server

import (
	lightweight "business_master/pb"
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Lightweight struct {
	service *grpc.Server
	listen  net.Listener
}

// server继承自动生成的服务类
type server struct {
	lightweight.UnimplementedGymServer
}

func Server_lightweight(listen net.Listener) {
	// 创建一个grpc服务器
	s := grpc.NewServer()
	l := &Lightweight{
		service: s,
		listen:  listen,
	}
	RegisterGymServer(l)
}

// 实现接口
func (s *server) BodyBuilding(ctx context.Context, in *lightweight.Person) (*lightweight.Reply, error) {
	fmt.Println("正在健身 ， 动作 ", in.Name, in.Actions)
	return &lightweight.Reply{Code: 0, Msg: "ok"}, nil
}

func RegisterGymServer(l *Lightweight) {
	// 注册服务
	lightweight.RegisterGymServer(l.service, &server{})
	if err := l.service.Serve(l.listen); err != nil {
		log.Error("failed to serve: ", err)
	}
}
