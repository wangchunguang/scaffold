package main

import (
	"business_master/grpc/server"
	log "github.com/sirupsen/logrus"
	"net"
)

const (
	grpc_port = ":8972"
)

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", grpc_port)
	if err != nil {
		log.Error("failed to listen :", err)
		return
	}
	// grpc与etcd 组合使用
	server.Server_user(listen)
	// 单独grpc连接使用
	//server.Server_lightweight(listen)
}
