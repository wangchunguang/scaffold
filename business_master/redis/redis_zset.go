package redis

import (
	"github.com/go-redis/redis/v7"
	log "github.com/sirupsen/logrus"
)

//  有序集合的操作 之所以有序是因为多了一个Score 参数
//添加一个或者多个元素到集合，如果元素已经存在则更新分数

func (r *RedisClient) ZAdd(key string, value ...*redis.Z) error {
	err := r.client.ZAdd(key, value...).Err()
	if err != nil {
		log.Error("redis ZAdd error --->", err)
		return err
	}
	return nil
}

// 查询元素对应的分数
func (r *RedisClient) ZScore(key, value string) (float64, error) {
	result, err := r.client.ZScore(key, value).Result()
	if err != nil {
		log.Error("redis ZScore error -->", err)
		return 0, err
	}
	return result, nil
}

// 根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
func (r *RedisClient) ZRank(key, value string) (int64, error) {
	result, err := r.client.ZRank(key, value).Result()
	if err != nil {
		log.Error("redis ZRank error --->", err)
		return 0, err
	}
	return result, nil
}

// 删除集合元素
func (r *RedisClient) ZRem(key string, value ...interface{}) error {
	err := r.client.ZRem(key, value).Err()
	if err != nil {
		log.Error("redis ZRem error--->", err)
		return err
	}
	return nil
}

// 获取元素中的个数
func (r *RedisClient) ZCard(key string) (int64, error) {
	result, err := r.client.ZCard(key).Result()
	if err != nil {
		log.Error("redis ZCard error--->", err)
		return 0, err
	}
	return result, nil
}

// 增加元素的分数
func (r *RedisClient) ZIncrBy(key, value string, num float64) error {
	err := r.client.ZIncrBy(key, num, value).Err()
	if err != nil {
		log.Error("redis ZIncrBy error --->", err)
		return err
	}
	return nil
}

// 返回集合中某个索引范围的元素，根据分数从小到大排序 ZRevRange 为从大到小
// start =0 stop =-1 表示全部
func (r *RedisClient) ZRange(key string, start, stop int64) ([]string, error) {
	result, err := r.client.ZRange(key, start, stop).Result()
	if err != nil {
		log.Error("redis ZRange error --->", err)
		return nil, err
	}
	return result, nil
}

// 根据索引范围删除元素
// 位置参数写成负数，代表从高分开始删除。
func (r *RedisClient) ZRemRangeByRank(key string, start, stop int64) (int64, error) {
	result, err := r.client.ZRemRangeByRank(key, start, stop).Result()
	if err != nil {
		log.Error("redis ZRemRangeByRank error --->", err)
		return result, err
	}
	return result, nil
}

//根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
// 初始化查询条件， Offset和Count用于分页
/*op := redis.ZRangeBy{
Min:"2", // 最小分数
Max:"10", // 最大分数
Offset:0, // 类似sql的limit, 表示开始偏移量
Count:5, // 一次返回多少数据
}*/
func (r *RedisClient) ZRangeByScore(key string, op redis.ZRangeBy) ([]string, error) {
	result, err := r.client.ZRangeByScore(key, &op).Result()
	if err != nil {
		log.Error("redis ZRangeByScore error --->", err)
		return nil, err
	}
	return result, nil
}

// 获取某个分段范围的元素个数 0 -1表示全部
//如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
func (r *RedisClient) ZCount(key, min, max string) (int64, error) {
	result, err := r.client.ZCount(key, min, max).Result()
	if err != nil {
		log.Error("redis ZCount error --->", err)
		return 0, err
	}
	return result, nil
}

// 根据分数范围删除元素
// 如果加上（ 则表示大于或者小于，相当于去掉了等于关系。
func (r *RedisClient) ZRemRangeByScore(key, min, max string) (int64, error) {
	result, err := r.client.ZRemRangeByScore(key, min, max).Result()
	if err != nil {
		log.Error("redis ZRemRangeByScore error --->", err)
		return 0, err
	}
	return result, nil
}
