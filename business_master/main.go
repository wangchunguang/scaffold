package main

import (
	"business_master/config"
	"business_master/dao"
	"business_master/redis"
	"flag"
	"fmt"
	"os"
	"os/signal"
)

var yaml_path = flag.String("y", "conf/conf_dev.yaml", "run yaml")

func main() {

	// 解析yaml
	yamlConfig, err := config.ReadYamlConfig(*yaml_path)
	if err != nil {
		return
	}
	// 连接mysql
	dao.MysqlInit(yamlConfig)
	// 连接redis
	redis.RedisInit(yamlConfig)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs)
	select {
	case s := <-sigs:
		// 这里做各种数据库的close操作
		fmt.Println(s)
	}

}
