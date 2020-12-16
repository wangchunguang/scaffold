package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

const (
	MQRETRYTIME = 3
)

// Yaml线上环境配置
type YamlConfig struct {
	Path          Path          `yaml:"path"`
	Mysql         Mysql         `yaml:"dao"`
	Redis         Redis         `yaml:"redis"`
	MongoDB       MongoDB       `yaml:"mongodb"`
	ElasticSearch ElasticSearch `yaml:"elastic_search"`
	RocketMQ      RocketMQ      `yaml:"rocket_mq"`
}

// 所有的文件配置路径
type Path struct {
	// 存储日志路径
	LogPath string `yaml:"log_path"`
}
type Mysql struct {
	//	用户名
	User string `yaml:"user"`
	//	密码
	Password string `yaml:"password"`
	//	ip地址
	Host string `yaml:"host"`
	//	端口号
	Port string `yaml:"port"`
	//	数据库名称
	Name string `yaml:"name"`
}
type Redis struct {
	//	用户名
	User string `yaml:"user"`
	//	密码
	Password string `yaml:"password"`
	//	ip地址
	Host string `yaml:"host"`
	//	端口号
	Port string `yaml:"port"`
	//	数据库名称
	DBNum string `yaml:"db_num"`
}

type MongoDB struct {
	//	用户名
	User string `yaml:"user"`
	//	密码
	Password string `yaml:"password"`
	//	ip地址
	Host string `yaml:"host"`
	//	端口号
	Port string `yaml:"port"`
	// 数据库
	DBName string `yaml:"db_name"`
	// 连接池数量
	MaxPoolSize string `yaml:"max_pool_size"`
}

type ElasticSearch struct {
	//	用户名
	User string `yaml:"user"`
	//	密码
	Password string `yaml:"password"`
	//	ip地址
	Host string `yaml:"host"`
}

type RocketMQ struct {
	NameServer []string `yaml:"name_server"`
}

// 解析yaml参数
func ReadYamlConfig(path string) (*YamlConfig, error) {
	yamlDev := &YamlConfig{}
	if f, err := os.Open(path); err != nil {
		log.Error(err)
		return nil, err
	} else {
		yaml.NewDecoder(f).Decode(yamlDev)
	}
	return yamlDev, nil

}
