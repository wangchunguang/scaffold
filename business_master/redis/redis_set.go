package redis

import (
	log "github.com/sirupsen/logrus"
)

func (r *RedisClient) SAdd(key string, value ...interface{}) error {
	err := r.client.SAdd(key, value).Err()
	if err != nil {
		log.Error("redis SAdd error --->", err)
		return err
	}
	return nil
}

// 获取元素的个数
func (r *RedisClient) SCard(key string) (int64, error) {
	result, err := r.client.SCard(key).Result()
	if err != nil {
		log.Error("redis SCard error --->", err)
		return 0, err
	}
	return result, nil
}

// 判断元素是否在集合中
func (r *RedisClient) SIsMember(key string, value interface{}) (bool, error) {
	result, err := r.client.SIsMember(key, value).Result()
	if err != nil {
		log.Error("redis SIsMember error --> ", err)
		return result, err
	}
	return result, nil
}

// 获取所有的元素
func (r *RedisClient) SMembers(key string) ([]string, error) {
	result, err := r.client.SMembers(key).Result()
	if err != nil {
		log.Error("redis SMembers error --->", err)
		return nil, err
	}
	return result, nil
}

// 删除元素
func (r *RedisClient) SRem(key string, value ...interface{}) error {
	err := r.client.SRem(key, value...).Err()
	if err != nil {
		log.Error("redis SRem error --->", err)
		return err
	}
	return nil
}

// 随机返回集合中的元素，并删除返回的元素,num代表返回元素的个数,没有则为0
func (r *RedisClient) SPop(key string, num int64) ([]string, error) {
	var values []string
	if num <= 0 {
		result, err := r.client.SPop(key).Result()
		if err != nil {
			log.Error("redis SPop error --->", err)
			return nil, err
		}
		values = append(values, result)
		return values, nil
	}
	result, err := r.client.SPopN(key, num).Result()
	if err != nil {
		log.Error("redis SPopN error --->", err)
		return nil, err
	}
	return result, nil
}
