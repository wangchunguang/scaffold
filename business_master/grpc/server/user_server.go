package server

import (
	"business_master/etcd"
	lightweight "business_master/pb"
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	etcd_ip = "127.0.0.1:2379"
	ip      = "127.0.0.1:8972"
	name    = "user_etcd"
)

type service struct {
	lightweight.UnimplementedUserServiceServer
}

// 新增邮件
func (s *service) SendMail(ctx context.Context, req *lightweight.MailRequest) (res *lightweight.MailResponse, err error) {
	if len(req.Mail) > 0 && len(req.Text) > 0 {
		fmt.Println("邮箱：", req.Mail, " 发送内容：", req.Text)
		return &lightweight.MailResponse{
			Ok: true,
		}, nil
	}
	return &lightweight.MailResponse{
		Ok: false,
	}, errors.New("No mailbox data is obtained")

}

//  获取用户信息
func (s *service) GetUserInfo(ctx context.Context, req *lightweight.User) (res *lightweight.MailReply, err error) {
	if req.Id < 0 {
		return &lightweight.MailReply{
			Code: 400,
			Msg:  "用户id不合法",
		}, errors.New("Invalid user id")
	}
	// 获取用户信息
	fmt.Println("id:", req.Id, " name:", req.Name, " age:", req.Age, " phone:", req.Phone, " sex:", req.Sex, " addr:", req.Addr)
	return &lightweight.MailReply{
		Code: 200,
		Msg:  "操作成功",
	}, nil
}

func Server_user(listen net.Listener) {
	// 创建服务器
	newServer := grpc.NewServer()
	//	 在grpc中注册服务
	lightweight.RegisterUserServiceServer(newServer, &service{})
	//	在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(newServer)
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	//	etcd服务注册:
	reg, err := etcd.NewService(etcd.ServiceInfo{
		Name: name,
		IP:   ip, // grpc节点
	}, []string{etcd_ip})
	if err != nil {
		log.Error(err)
	}
	go reg.Start()
	if err := newServer.Serve(listen); err != nil {
		log.Error("server listen err:", err)
	}
}
