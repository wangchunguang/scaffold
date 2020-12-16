package mongoDB

import (
	"business_master/config"
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"time"
)

var (
	mongoClient *mongo.Client
	err         error
)

func MongoInit(yamlConfig *config.YamlConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	num, _ := strconv.Atoi(yamlConfig.MongoDB.MaxPoolSize)
	if mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+yamlConfig.MongoDB.Host+":"+yamlConfig.MongoDB.Port).SetMaxPoolSize(uint64(num))); err != nil {
		log.Error("mongodb client error ", err)
		return
	}
	// 检查连接
	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Error("mongodb ping error ", err)
		return
	}
	log.Info("mongodb  connection succeeded")

}
