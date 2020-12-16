package rocket_mq

import (
	"business_master/config"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	log "github.com/sirupsen/logrus"
)

var (
	mq_name  = "hcg_wc"
	RocketMQ *RocketMqClient
)

type RocketDriver struct {
	NameServer []string
}

func rocketmq_init(mq *RocketDriver) {
	if mq == nil {
		log.Error("mq config is nil")
		return
	}
	// 创建消费者
	MqConsumer, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(mq_name),
		consumer.WithNameServer(mq.NameServer))
	if err != nil {
		log.Error("init rocket mq push consumer err: ", err)
		return
	}

	// 创建生产者
	MqProducer, err := rocketmq.NewProducer(
		// 地址
		producer.WithNameServer(mq.NameServer),
		// 名字服务名称
		producer.WithGroupName(mq_name),
		// 重试次数
		producer.WithRetry(config.MQRETRYTIME),
		// 主题命名空间
		//producer.WithNamespace("namespace"),
	)
	if err != nil {
		log.Error("rocket_mq connection failure")
		return
	}
	err = MqProducer.Start()
	if err != nil {
		log.Error("start producer error: ", err)
		return
	}

	RocketMQ = NewRocketMqClient(MqProducer, MqConsumer)
}

func RocketInit(yamlConfig *config.YamlConfig) {
	mq := NewRocketMQ(yamlConfig)
	rocketmq_init(mq)
}

func NewRocketMQ(yamlConfig *config.YamlConfig) *RocketDriver {
	return &RocketDriver{
		NameServer: yamlConfig.RocketMQ.NameServer,
	}
}

func ShutDownMQ(p *RocketMqClient) {
	p.Producer.Shutdown()
	p.Consumer.Shutdown()
}
