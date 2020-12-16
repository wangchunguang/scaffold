package rocket_mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	log "github.com/sirupsen/logrus"
	"os"
	"sync"
	"time"
)

var sy_once sync.Once

type RocketMqClient struct {
	Producer rocketmq.Producer
	Consumer rocketmq.PushConsumer
}

func NewRocketMqClient(MqProducer rocketmq.Producer, MqConsumer rocketmq.PushConsumer) *RocketMqClient {
	return &RocketMqClient{
		Producer: MqProducer,
		Consumer: MqConsumer,
	}
}

// 订阅
func (p *RocketMqClient) ConsumeSuccess(topic string) {

	err := p.Consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			log.Info(string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		log.Error("consumer subscribe TopicNotifySucess err =", err)
		return
	}
	once_start := func() {
		err := p.Consumer.Start()
		if err != nil {
			log.Error("rocket consume start err ", err)
			os.Exit(-1)
			return
		}
	}
	sy_once.Do(once_start)

	//time.Sleep(time.Millisecond*10)

}

// 退订
func (p *RocketMqClient) Unsubscribe(topic string) error {
	err := p.Consumer.Unsubscribe(topic)
	if err != nil {
		log.Error("rocket_mq Unsubscribe error :", err)
		return err
	}
	return nil
}

// 发布

// 同步发送
func (p *RocketMqClient) ProducerSendSync(msg *primitive.Message) (*primitive.SendResult, error) {
	sync, err := p.Producer.SendSync(context.Background(), msg)
	if err != nil {
		log.Error("send message error :", err)
		return nil, err
	}
	return sync, nil
}

// 异步发送
func (p *RocketMqClient) ProducerSendASync(msg *primitive.Message) error {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	err := p.Producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		if err != nil {
			log.Error("receive message error: ", err)
		} else {
			log.Info("send message success: result=", result.String())
		}
		wg.Done()
	}, msg)
	if err != nil {
		log.Error("send message error:", err)
		return err
	}
	return nil
}

// 延时发送
func (p *RocketMqClient) ProducerSendDelay(msg *primitive.Message, sleepTime int64) (*primitive.SendResult, error) {
	time.Sleep(time.Duration(sleepTime))
	sync, err := p.Producer.SendSync(context.Background(), msg)
	if err != nil {
		log.Error("send message error :", err)
		return nil, err
	}
	return sync, nil
}
