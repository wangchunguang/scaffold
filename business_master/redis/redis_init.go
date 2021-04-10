package redis

import (
	"business_master/config"
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func RedisInit(yamlConfig *config.YamlConfig) *redis.Client {
	db, _ := strconv.Atoi(yamlConfig.Redis.DBNum)
	client := redis.NewClient(&redis.Options{
		Addr:         yamlConfig.Redis.Host + ":" + yamlConfig.Redis.Port,
		Password:     yamlConfig.Redis.Password,
		DB:           db,
		MaxRetries:   5,
		MinIdleConns: 5,
		PoolSize:     5,
	}).WithTimeout(5 * time.Second)
	result, err := client.Ping().Result()
	if err != nil {
		log.Error("redis client error ", err)
		return nil
	}
	log.Info("redis  connection succeeded ", result)
	return client
}

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(yaml *config.YamlConfig) *RedisClient {
	return &RedisClient{
		RedisInit(yaml),
	}

}

func InterfaceTOString(result []interface{}) []string {
	values := make([]string, len(result))
	for k, v := range result {
		values[k] = v.(string)
	}
	log.Info("InterfaceTOString .... ", values)
	return values
}
