package redis

import (
	"business_master/config"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"testing"
)

func Test_Redis(t *testing.T) {

	yamlConfig := &config.YamlConfig{
		Redis: config.Redis{
			User:     "",
			Password: "",
			Host:     "127.0.0.1",
			Port:     "6379",
			DBNum:    "1",
		},
	}
	client := NewRedisClient(yamlConfig)
	pubSub(client)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	select {
	case s := <-sigs:
		// 这里做各种数据库的close操作
		fmt.Println(s)
	}
}

func pubSub(client *RedisClient) {
	ch := make(chan string, 1)
	go pubdemo(client)
	go client.Subscribe("wang", ch)
	for {
		select {
		case value := <-ch:
			fmt.Println("管道的数据------->", value)
		}
	}
}

func pubdemo(client *RedisClient) {
	for i := 1; i <= 10; i++ {
		client.Publish("wang", strconv.Itoa(i))
	}
}
