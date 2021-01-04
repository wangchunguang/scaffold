package main

import (
	"github.com/prometheus/common/log"
	"master/scaffold_master/grpc/client"
	"os"
	"os/signal"
)

const (
	address = "localhost:50051"
)

func main() {
	// grpc 单独使用
	/*	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
		defer conn.Close()
		if err != nil {
			log.Error("did not connect:", err)
			return
		}

		client.Client_lightweight(conn)*/

	client.Client_user()
	s := make(chan os.Signal)
	signal.Notify(s)
	select {
	case c := <-s:
		log.Info("Exit  ", c)
	}
}
