package elasticSearch

import (
	"business_master/config"
	"context"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

type Elastic struct {
	esClient *elastic.Client
}

func NewElastic(yaml *config.YamlConfig) *Elastic {
	return &Elastic{ElasticSearchInit(yaml)}
}

// 查询excel是否存在
func (e *Elastic) IndexExists(_index ...string) (bool, error) {
	exists, err := e.esClient.IndexExists(_index...).
		Do(context.Background())
	if err != nil {
		log.Error("es index exists error ", err)
		return false, err
	}
	return exists, nil
}

//创建 index
func (e *Elastic) CreateIndex(_index, mapping string) (bool, error) {
	result, err := e.esClient.CreateIndex(_index).
		BodyString(mapping).
		Do(context.Background())
	if err != nil {
		log.Error("es create index error ", err)
		return false, err
	}
	return result.Acknowledged, nil
}

// 删除index
func (e *Elastic) EsDeleteIndex(index ...string) error {
	_, err := e.esClient.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		log.Error("es delete index error ", err)
		return err
	}
	return nil
}

// 新增一条数据
func (e *Elastic) EsSave(_index, _id string, value interface{}) (*elastic.IndexResponse, error) {
	response, err := e.esClient.Index().
		Index(_index).
		Id(_id).BodyJson(value).
		Do(context.Background())
	if err != nil {
		log.Error("es save error ", err)
		return nil, err
	}
	return response, nil
}

// 通过id查询一条数据
func (e *Elastic) EsSelectById(_index, _id string) ([]byte, error) {
	response, err := e.esClient.Get().
		Index(_index).
		Id(_id).
		Do(context.Background())
	if err != nil {
		log.Error("es select by id error ", err)
		return nil, err
	}
	json, err := response.Source.MarshalJSON()
	if err != nil {
		log.Error("")

	}
	return json, nil
}

// 批量查找文档
// boolQuery := elastic.NewBoolQuery().Must() 可以组装多个条件
//Aggregation 是设置聚合条件
func (e *Elastic) EsSelectList(_index string, boolQuery *elastic.BoolQuery, start, size int) (*elastic.SearchResult, error) {
	result, err := e.esClient.Search().
		Index(_index).
		Query(boolQuery).
		From(start).
		Size(size).
		Do(context.Background())
	if err != nil {
		log.Error("es select list error :", err)
		return nil, err
	}
	return result, nil
}

// 根据id更新文档
func (e *Elastic) EsUpdateId(_index, _id string, doc interface{}) error {
	_, err := e.esClient.
		Update().
		Index(_index).
		Id(_id).
		Doc(doc).
		Do(context.Background())
	if err != nil {
		log.Error("es update id =", _id, " error ", err)
		return err
	}
	return nil
}

// 批量更新	elastic.BoolQuery.Must(termQuery,termQuerys)
func (e *Elastic) EsUpdateByQuery(_index, _source string, query *elastic.BoolQuery) error {
	_, err := e.esClient.
		UpdateByQuery(_index).
		Query(query).
		Script(elastic.NewScript(_source)).
		ProceedOnVersionConflict().
		Do(context.Background())
	if err != nil {
		log.Error("es update by query error ", err)
		return err
	}
	return nil
}

// 根据id删除文档
func (e *Elastic) EsDeleteById(_index, _id string) error {
	_, err := e.esClient.Delete().
		Index(_index).
		Id(_id).
		Do(context.Background())
	if err != nil {
		log.Error("es delete id error", err)
		return err
	}
	return nil
}

// 通过条件删除文档
func (e *Elastic) EsDeleteByQuery(_index string, query *elastic.BoolQuery) error {
	_, err := e.esClient.DeleteByQuery(_index).
		Query(query).
		ProceedOnVersionConflict().
		Do(context.Background())
	if err != nil {
		log.Error("es delete by query error ", err)
		return err
	}
	return nil
}
