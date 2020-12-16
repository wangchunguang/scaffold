package mongoDB

import (
	"business_master/config"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
	"time"
)

type TimeTest struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}
type Record struct {
	JobName string   `bson:"jobName"` //任务名
	Command string   `bson:"command"` //shell命令
	Err     string   `bson:"err"`     //脚本错误
	Content string   `bson:"content"` //脚本输出
	Tp      TimeTest //执行时间
}

func Test_MongoDB(t *testing.T) {
	yamlConfig := &config.YamlConfig{
		/*		Path: config.Path{
				"\\logs\\InputLogs\\",
			},*/
		MongoDB: config.MongoDB{
			User:        "",
			Password:    "",
			Host:        "127.0.0.1",
			Port:        "27017",
			DBName:      "demo",
			MaxPoolSize: "20",
		},
	}
	// 开启日志
	//logs.InitLog(yamlConfig.Path.LogPath)
	MongoInit(yamlConfig)
	// 新增一条数据
	//test_setInsertOne()
	// 新增多条数据
	//test_insertMany()
	//	 查询单个数据
	//	test_selectFindOne()
	//	根据条件查询数据
	//test_selectFind()
	//	 查询集合里面有多少数据
	//test_collectionCount()
	//	 分页查询
	//test_CollectionDocuments()
	// 修改一条数据
	//test_update_one()
	// 批量修改
	//test_update_many()
	//	 根据条件删除一个
	//test_delete_one()
	// 根据条件删除多条数据
	//test_delete_many()
	//	传建一个索引
	//test_crate_index_one()
	// 创建多个索引
	//test_crate_index_many()
	// 获取全部索引
	//	test_select_list()
	// 删除一个索引
	//text_delete_one_index()
	// 删除多个索引
	//text_delete_many()
}

func NewMongoDriverInit() *MongoDriver {
	driver := NewMongoDriver("demo", "records")
	return driver
}

func test_setInsertOne() {
	record := &Record{
		JobName: "job1",
		Command: "command1",
		Err:     "not err",
		Content: "12",
		Tp: TimeTest{
			StartTime: time.Now().Unix(),
			EndTime:   time.Now().Unix(),
		},
	}
	driver := NewMongoDriverInit()
	driver.InsertOne(record)
}

func test_insertMany() {
	records := []interface{}{
		Record{
			JobName: "job1",
			Command: "echo 1",
			Err:     "error 1",
			Content: "1",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job2",
			Command: "echo 2",
			Err:     "error 2",
			Content: "2",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job3",
			Command: "echo 3",
			Err:     "error 3",
			Content: "3",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job4",
			Command: "echo 4",
			Err:     "error 4",
			Content: "4",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job5",
			Command: "echo 5",
			Err:     "error 5",
			Content: "5",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job6",
			Command: "echo 6",
			Err:     "error 6",
			Content: "6",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		Record{
			JobName: "job2",
			Command: "echo 2",
			Err:     "error 2",
			Content: "2",
			Tp: TimeTest{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	}
	driver := NewMongoDriverInit()
	driver.InsertMany(records)

}

func test_selectFindOne() {
	var record *Record
	driver := NewMongoDriverInit()
	bson_one := bson.M{"jobName": "job2", "content": bson.M{"$lte": "5"}}
	one, _ := driver.FindOne(bson_one)
	bytes, err := one.DecodeBytes()
	if err != nil {
		fmt.Println(err)
	}
	err = bson.Unmarshal(bytes, &record)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(record)
}

func test_selectFind() {
	var results []Record
	driver := NewMongoDriverInit()
	//	bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
	// and
	//bson := bson.D{{"jobName", "job2"}, {"content", "5"}}
	// 不等于 !=($ne)
	//bson := bson.M{"jobName": "job2", "content":bson.M{"$ne":"5"}}
	// 大于>($gt)
	//bson := bson.M{"jobName": "job2", "content":bson.M{"$gt":"2"}}
	// 小于<($lt)
	//bson := bson.M{"jobName": "job2", "content":bson.M{"$lt":"5"}}
	// 大于等于>=($gte)
	//bson := bson.M{"jobName": "job2", "content":bson.M{"$gte":"5"}}
	// 小于等于<=($lte)
	bson := bson.M{"jobName": "job2", "content": bson.M{"$lte": "5"}}
	// no in($nin)同$in
	// 是否包含这个键($exists)
	// 正则匹配($regex)
	// $or 或者
	// 修改$set
	// 增加值$incr
	cursor, err := driver.Find(bson)
	if err != nil {
		fmt.Println("err  Find", err)
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		fmt.Println("error cursor all ", err)
	}
	marshal, _ := json.Marshal(results)
	fmt.Println(string(marshal))
}

func test_collectionCount() {
	driverInit := NewMongoDriverInit()
	name, count, err := driverInit.CollectionCount()
	if err != nil {
		fmt.Println(" collectionCount error ", err)
	}
	fmt.Println(name)
	fmt.Println(count)
}

func test_CollectionDocuments() {
	var record []Record
	driver := NewMongoDriverInit()
	documents, err := driver.CollectionDocuments(1, 5, 1, nil)
	if err != nil {
		fmt.Println("CollectionDocuments error ", err)
	}
	if err := documents.All(context.TODO(), &record); err != nil {
		fmt.Println("CollectionDocuments all error ", err)
	}
	marshal, _ := json.Marshal(record)
	fmt.Println(string(marshal))
}

func test_update_one() {
	driver := NewMongoDriverInit()
	filter := bson.M{"jobName": "job2"}
	update := bson.M{"$set": bson.M{"err": "11111"}}
	_, err := driver.UpdateOne(filter, update)
	if err != nil {
		fmt.Println("update_one error ", err)
	}
}

func test_update_many() {
	filter := bson.M{"jobName": "job2"}
	update := bson.M{"$set": bson.M{"err": "22222"}}
	many, _ := NewMongoDriverInit().UpdateMany(filter, update)
	fmt.Println(many.ModifiedCount)
}

func test_delete_one() {

	NewMongoDriverInit().DeleteOne(bson.M{"jobName": "job2"})

}

func test_delete_many() {
	many, _ := NewMongoDriverInit().DeleteMany(bson.M{"jobName": "job2"})
	fmt.Println(many)
}

func test_crate_index_one() {
	model := mongo.IndexModel{
		Keys: bson.M{"err": -1},
	}
	one, _ := NewMongoDriverInit().CreatIndexOne(model)
	fmt.Println(one)
}

func test_crate_index_many() {
	model := []mongo.IndexModel{
		{
			Keys: bson.D{{"jobName", "text"}, {"content", 1}},
		},
		{
			Keys: bson.D{{"content", 1}},
		},
	}
	one, _ := NewMongoDriverInit().CreatIndexMany(model)
	fmt.Println(one)
}

func test_select_list() {
	list, _ := NewMongoDriverInit().SelectIndexList()
	defer list.Close(context.TODO())
	for list.Next(context.TODO()) {
		keyElem, _ := list.Current.LookupErr("key")
		fmt.Println(keyElem)
	}
}

func text_delete_one_index() {
	index, _ := NewMongoDriverInit().DeleteOneIndex("command_1")
	fmt.Println(index)
}

func text_delete_many() {
	many, _ := NewMongoDriverInit().DeleteIndexMany()
	fmt.Println(many)
}
