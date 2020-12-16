package redis

import (
	log "github.com/sirupsen/logrus"
)

// 发布消息
func (r *RedisClient) Publish(channel, message string) error {

	err := r.client.Publish(channel, message).Err()
	if err != nil {
		log.Error("redis push error --->", err)
		return err
	}
	log.Info("redis Publish message--->", message)
	return nil
}

// 订阅一个频道 并接收消息
func (r *RedisClient) Subscribe(channel string, ch chan string) error {
	subscribe := r.client.Subscribe(channel)
	_, err := subscribe.Receive()
	if err != nil {
		log.Error("redis Subscribe error --->", err)
		return nil
	}
	// 订阅过后返回的消息
	messages := subscribe.Channel()
	for value := range messages {
		log.Info("redis value.Payload --->", value.Payload)
		ch <- value.Payload
	}
	return nil
}

// 查询指定的channel上面有多少订阅者
func (r *RedisClient) PubSubNumSub(channel ...string) (map[string]int64, error) {
	result, err := r.client.PubSubNumSub(channel...).Result()
	if err != nil {
		log.Error("redis PubSubNumSub error --->", err)
		return nil, err
	}
	return result, nil
}
