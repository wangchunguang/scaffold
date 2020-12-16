package redis

import log "github.com/sirupsen/logrus"

// 左边插入
func (r *RedisClient) LPush(key string, value ...interface{}) error {
	err := r.client.LPush(key, value).Err()
	if err != nil {
		log.Error("redis LPush error --->", err)
		return err
	}
	return nil
}

// 删除左边第一个参数 并返回
func (r *RedisClient) LPop(key string) (string, error) {
	result, err := r.client.LPop(key).Result()
	if err != nil {
		log.Error("redis LPop error --->", err)
		return "", err
	}
	return result, nil
}

// 删除右边的第一个数据 并返回
func (r *RedisClient) RPop(key string) (string, error) {
	result, err := r.client.RPop(key).Result()
	if err != nil {
		log.Error("redis RPop error --->", err)
		return "", err
	}
	return result, err
}

// 从右边插入一个数据
func (r *RedisClient) RPush(key string, value ...interface{}) (int64, error) {
	result, err := r.client.RPush(key, value).Result()
	if err != nil {
		log.Error("redis RPush error --->", err)
		return 0, err
	}
	return result, nil
}

// 返回列表的大小
func (r *RedisClient) LLen(key string) (int64, error) {
	result, err := r.client.LLen(key).Result()
	if err != nil {
		log.Error("redis LLen error --->", err)
		return 0, err
	}
	return result, nil
}

// 返回一个返回的数据 end为-1 表示返回当前截止最后的一个数据
func (r *RedisClient) LRange(key string, start, end int64) ([]string, error) {
	result, err := r.client.LRange(key, start, end).Result()
	if err != nil {
		log.Error("redis LRange error --->", err)
		return nil, err
	}
	return result, nil
}

// 删除列表中重复的数据 从列表左边开始，如果出现重复的数，根据删除num次
func (r *RedisClient) LRem(key, value string, num int64) (int64, error) {
	result, err := r.client.LRem(key, num, value).Result()

	if err != nil {
		log.Error("redis LRem error --->", err)
		return 0, err
	}
	return result, nil
}

// 根据索查询数据
func (r *RedisClient) LIndex(key string, num int64) (string, error) {
	result, err := r.client.LIndex(key, num).Result()
	if err != nil {
		log.Error("redis LIndex error --->", err)
		return "", err
	}
	return result, nil
}

// 在指定位置插入数据  value 表示插入的值 original表示索引位置的值
//numType 为1表示之前插入 2表示之后插入
func (r *RedisClient) LInsert(key, value, original string, numType int64) error {
	if numType == 1 {
		err := r.client.LInsert(key, "before", original, value).Err()
		if err != nil {
			log.Error("redis before LInsert error --->", err)
			return err
		}
		return nil

	} else if numType == 2 {
		err := r.client.LInsert(key, "after", original, value).Err()
		if err != nil {
			log.Error("redis after LInsert error --->", err)
			return err
		}
		return nil
	}
	return nil
}
