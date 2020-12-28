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

type service struct {
	lightweight.UnimplementedUserServiceServer
}

// 新增邮件
func (s *service) SendMail(ctx context.Context, req *lightweight.MailRequest) (res *lightweight.MailResponse, err error) {
	if len(req.Mail) > 0 && len(req.Text) > 0 {
		fmt.Printf("邮箱:%s;发送内容:%s", req.Mail, req.Text)
		res.Ok = true
		return res, nil
	}
	res.Ok = false
	log.Error("No mailbox data is obtained")
	return res, errors.New("No mailbox data is obtained")
}

//  获取用户信息
func (s *service) GetUserInfo(ctx context.Context, req *lightweight.User) (res *lightweight.MailReply, err error) {
	if req.Id < 0 {
		res.Code = 400
		res.Msg = "用户id不合法"
		return res, errors.New("Invalid user id")
	}
	// 获取用户信息
	fmt.Println(req)
	res.Code = 200
	return res, nil
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
	//	etcd服务注册
	reg, err := etcd.NewService(etcd.ServiceInfo{
		Name: "user",
		IP:   "127.0.0.1:8972", // grpc节点
	}, []string{"127.0.0.1:2379"})
	if err != nil {
		log.Error(err)
	}
	go reg.Start()
	if err := newServer.Serve(listen); err != nil {
		log.Println(err)
	}
}
