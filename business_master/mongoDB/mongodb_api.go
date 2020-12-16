package mongoDB

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriver struct {
	database   string
	collection string
}

func NewMongoDriver(database, collection string) *MongoDriver {
	return &MongoDriver{
		database,
		collection,
	}
}

// 插入单个数据
func (m *MongoDriver) InsertOne(value interface{}) (*mongo.InsertOneResult, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	insertResult, err := collection.InsertOne(context.TODO(), value)
	if err != nil {
		log.Error("mongodb InsertOne error ", err)
		return nil, err
	}
	return insertResult, nil
}

// 插入多个数据
func (m *MongoDriver) InsertMany(value []interface{}) (*mongo.InsertManyResult, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	many, err := collection.InsertMany(context.TODO(), value)
	if err != nil {
		log.Error("mongodb InsertMany error ", err)
		return nil, err
	}
	return many, nil
}

// 根据条件查询一条数据
func (m *MongoDriver) FindOne(filter interface{}) (*mongo.SingleResult, error) {
	collection, err := mongoClient.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		log.Error("mongodb FindOne error", err)
		return nil, err
	}
	singleResult := collection.FindOne(context.TODO(), filter)
	return singleResult, nil
}

//  根据条件查询符合的数据
func (m *MongoDriver) Find(filter interface{}) (*mongo.Cursor, error) {
	collection, err := mongoClient.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		log.Error("mongodb Find error", err)
		return nil, err
	}
	find, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Error("mongodb select find error", err)
		return nil, err
	}
	return find, nil
}

//  查询集合有多少数据
func (m *MongoDriver) CollectionCount() (string, int64, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	name := collection.Name()
	count, err := collection.EstimatedDocumentCount(context.TODO())
	if err != nil {
		log.Error("mongodb CollectionCount error  ", err)
		return "", 0, err
	}
	return name, count, nil
}

// 按选项查询集合
// skip方法同样接受一个数字参数作为跳过的记录条数。（前N个不要）,默认值是0
// sort() 方法可以通过参数指定排序的字段，并使用 1 和 -1 来指定排序的方式，其中 1 为升序排列，而 -1 是用 于降序排列
// 每页查询多少个
// filter 条件
func (m *MongoDriver) CollectionDocuments(Skip, Limit int64, sort int, filter interface{}) (*mongo.Cursor, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	findOptions := options.Find().SetSort(sort).SetLimit(Limit).SetSkip(Skip)
	find, err := collection.Find(context.Background(), filter, findOptions)
	if err != nil {
		log.Error("mongodb CollectionDocuments error ", err)
		return nil, err
	}
	return find, nil
}

// 修改文档 修改一条数据
func (m *MongoDriver) UpdateOne(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Error("mongodb updateOne error ", err)
		return nil, err
	}
	return result, nil
}

// 修改修改文档 批量修改
func (m *MongoDriver) UpdateMany(filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	many, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		log.Error("mongodb updateMany error ", err)
		return nil, err
	}
	return many, nil
}

//根据条件删除一条数据
func (m *MongoDriver) DeleteOne(filter interface{}) (int64, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	if err != nil {
		log.Error("mongodb DeleteAndFind error ", err)
		return 0, err
	}
	return count.DeletedCount, nil
}

//根据条件删除多条数据
func (m *MongoDriver) DeleteMany(filter interface{}) (int64, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	count, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		log.Error("mongodb DeleteMany error ", err)
		return 0, err
	}
	return count.DeletedCount, nil
}

// 创建一个索引
func (m *MongoDriver) CreatIndexOne(model mongo.IndexModel) (string, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	index, err := collection.Indexes().CreateOne(context.TODO(), model)
	if err != nil {
		log.Error("mongodb creat index one error ", err)
		return "", err
	}
	return index, nil
}

// 创建多个索引
func (m *MongoDriver) CreatIndexMany(models []mongo.IndexModel) ([]string, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	many, err := collection.Indexes().CreateMany(context.TODO(), models)
	if err != nil {
		log.Error("mongodb creat index many error ", err)
		return nil, err
	}
	return many, nil
}

// 获取所有的索引
func (m *MongoDriver) SelectIndexList() (*mongo.Cursor, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	list, err := collection.Indexes().List(context.TODO())
	if err != nil {
		log.Error("mongodb select index list error ", err)
		return nil, err
	}
	return list, nil
}

// 删除单个索引
func (m *MongoDriver) DeleteOneIndex(name string) (bson.Raw, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	one, err := collection.Indexes().DropOne(context.TODO(), name)
	if err != nil {
		log.Error("mongodb delete index one error ", err)
		return nil, err
	}
	return one, nil
}

// 删除该集合的所有索引
func (m *MongoDriver) DeleteIndexMany() (bson.Raw, error) {
	collection := mongoClient.Database(m.database).Collection(m.collection)
	all, err := collection.Indexes().DropAll(context.TODO())
	if err != nil {
		log.Error("mongodb delete many index error ", err)
		return nil, err
	}
	return all, nil
}
