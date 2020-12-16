package elasticSearch

import (
	"business_master/config"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"strconv"
	"testing"
)

type Subject struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Num     int    `json:"num"`
	Content string `json:"content"`
	Genres  string `json:"genres"`
	Address string `json:"address"`
}

const (
	hc  = "hc"
	wcg = "wcg"
)

const mapping = `
			{
				"mappings": {
					"properties": {
						"id": {
							"type": "long"
						},
						"title": {
							"type": "text"
						},
						"genres": {
							"type": "keyword"
						}
					}
				}
			}`

func ESInit() *elastic.Client {
	yamlConfig := &config.YamlConfig{
		ElasticSearch: config.ElasticSearch{
			User:     "",
			Password: "",
			Host:     "http://127.0.0.1:9200",
		},
	}
	return ElasticSearchInit(yamlConfig)
}

func TestIndexExists(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	exists, err := el.IndexExists(wcg)
	if err != nil {
		log.Error("error IndexExists  ", err)
		return
	}
	if !exists {
		index, err := el.CreateIndex(wcg, mapping)
		if err != nil {
			return
		}
		log.Info("creat index ", index)
	}
}

func TestElastic_EsSave(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	su := &Subject{
		ID:      10,
		Title:   "金杨",
		Num:     18,
		Content: "轮滑",
		Genres:  "体育",
		Address: "四川成都",
	}
	el.EsSave(wcg, strconv.Itoa(su.ID), su)
}

func TestElastic_EsUpdateId(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	su := &Subject{
		ID:    7,
		Title: "红尘滚滚",
	}
	el.EsUpdateId(hc, strconv.Itoa(su.ID), su)
}

func TestElastic_EsUpdateByQuery(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	d1 := elastic.NewQueryStringQuery("title:游戏")
	d2 := elastic.NewQueryStringQuery("id:4")
	must := elastic.NewBoolQuery().Must(d1, d2)
	content := "666"
	//address := "上海"
	scrpt := "ctx._source['num']='" + content + "'"
	_ = el.EsUpdateByQuery(wcg, scrpt, must)
}

func TestElastic_EsDeleteById(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	el.EsDeleteById(hc, "1")
}

func TestElastic_EsDeleteByQuery(t *testing.T) {
	el := &Elastic{
		esClient: ESInit(),
	}
	d1 := elastic.NewMatchPhraseQuery("title", "红尘滚滚")
	d2 := elastic.NewMatchPhraseQuery("num", "666")
	query := elastic.NewBoolQuery().Must(d1, d2)
	el.EsDeleteByQuery(hc, query)

}
