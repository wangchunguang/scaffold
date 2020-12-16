package rocket_mq

import (
	"business_master/config"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"testing"
)

var (
	topic = "HC_WCG_MQ"
	tag   = "运动"
)

const (
	SendOK = iota
	SendFlushDiskTimeout
	SendFlushSlaveTimeout
	SendSlaveNotAvailable
	SendUnknownError
)

func Test_RocketMQ(t *testing.T) {
	yaml_config := &config.YamlConfig{
		RocketMQ: config.RocketMQ{
			NameServer: []string{"127.0.0.1:9876"},
		},
	}
	RocketInit(yaml_config)
}

func TestRocketMqClient_ProducerSendSync(t *testing.T) {
	Test_RocketMQ(t)
	go Consume_Test()
	go Producer_Test()
	sign := make(chan os.Signal, 1)
	signal.Notify(sign)
	select {
	case s := <-sign:
		ShutDownMQ(RocketMQ)
		fmt.Println(s)
	}

}

func Consume_Test() {
	RocketMQ.ConsumeSuccess(topic)
}

func Producer_Test() {
	for {
		body := "散步"
		msg := &primitive.Message{
			Topic:         topic,
			Body:          []byte(body),
			Flag:          0,
			TransactionId: "1",
			Batch:         false,
			Queue:         nil,
		}
		msg.WithTag(tag)
		sync, err := RocketMQ.ProducerSendSync(msg)
		if err != nil {
			log.Error("rocket_mq sync error ", err)
			return
		}
		status(int(sync.Status), sync.MsgID)
	}
}

func status(code int, msgId string) {
	switch code {
	case SendOK:
		log.Info("发送成功,msgId=", msgId)
	case SendFlushDiskTimeout:
		log.Info("发送超时")
	case SendFlushSlaveTimeout:
		log.Info("刷新磁盘超时")
	case SendSlaveNotAvailable:
		log.Info("发送站不可用")
	case SendUnknownError:
		log.Info("发送错误")
	}
}
