package client

import (
	"context"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	lightweight "master/scaffold_master/pb"
	"os"
	"time"
)

func Client_lightweight(conn *grpc.ClientConn) {
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
