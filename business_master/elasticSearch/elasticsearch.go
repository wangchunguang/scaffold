package elasticSearch

import (
	"business_master/config"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"time"
)

func ElasticSearchInit(yamlConfig *config.YamlConfig) *elastic.Client {
	esClient, err := elastic.NewClient(
		// 服务器地址，多个服务器地址用 ， 隔开
		elastic.SetURL(yamlConfig.ElasticSearch.Host),
		// 认证的账号和密码
		elastic.SetBasicAuth(yamlConfig.ElasticSearch.User, yamlConfig.ElasticSearch.Password),
		// 启动gzip压缩
		elastic.SetGzip(true),
		// 设置监控查询时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
	)
	if err != nil {
		log.Error("ElasticSearch client error --->", err)
		return nil
	}
	log.Info("ElasticSearch connection succeeded")
	return esClient
}
