package main

import (
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"os/signal"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()
	if err != nil {
		log.Error("did not connect:", err)
		return
	}
	// grpc 单独使用
	client.Client_lightweight(conn)
	s := make(chan os.Signal)
	signal.Notify(s)
	select {
	case c := <-s:
		log.Info("Exit  ", c)
	}
}
