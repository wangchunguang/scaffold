package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	lightweight "scaffold/pb"
	"time"
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
	client := lightweight.NewGymClient(conn)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	ch := make(chan lightweight.Person, 10)
	li := &lightweight.Person{
		Name:    "红尘",
		Actions: []string{"抽烟", "喝酒", "打牌"},
	}
	ch <- *li
	save_lightweight(ch, client, ctx)
	s := make(chan os.Signal)
	signal.Notify(s)
	select {
	case c := <-s:
		log.Info("Exit  ", c)
	}

}

func save_lightweight(ch chan lightweight.Person, client lightweight.GymClient, ctx context.Context) {
	for {
		select {
		case p := <-ch:
			building, err := client.BodyBuilding(ctx, &p)
			if err != nil {
				log.Error("err ", err)
				os.Exit(0)
				return
			}
			log.Info(" code=", building.Code, " msg=", building.Msg)
		default:
			time.Sleep(time.Second)
		}
	}
}
