package redis

import (
	log "github.com/sirupsen/logrus"
	"time"
)

// set
func (r *RedisClient) Set(key, value string, num time.Duration) error {
	var err error
	if num > 0 {
		err = r.client.Set(key, value, num).Err()
	} else {
		err = r.client.Set(key, value, 0).Err()
	}
	if err != nil {
		log.Info("redis set error -->", err)
		return err
	}
	return nil
}

// 批量设置 Mset
func (r *RedisClient) MSet(args ...interface{}) error {
	err := r.client.MSet(args...).Err()
	if err != nil {
		log.Error("redis MSet  err", err)
		return err
	}
	return nil
}

// get
func (r *RedisClient) StringGet(key string) (string, error) {
	result, err := r.client.Get(key).Result()
	if err != nil {
		log.Error("redis get  error ----->", err)
		return "", err
	}
	return result, nil
}

// 批量查询key的值
func (r *RedisClient) MGet(args []string) ([]string, error) {

	result, err := r.client.MGet(args...).Result()
	if err != nil {
		log.Error("redis MGet error  ", err)
		return nil, err
	}
	return InterfaceTOString(result), nil
}

// 返回key的旧值
func (r *RedisClient) GetSet(key, value string) (string, error) {
	result, err := r.client.GetSet(key, value).Result()
	if err != nil {
		log.Error("redis get set error----> ", err)
		return "", err
	}
	return result, nil
}

// 设置nx 简单的分布式锁的使用
func (r *RedisClient) SetNx(key, value string, num time.Duration) error {
	err := r.client.SetNX(key, value, num).Err()
	if err != nil {
		log.Error("redis setNx err---->", err)
		return err
	}
	return nil
}

// 自增加一
func (r *RedisClient) Incr(key string) (int64, error) {
	result, err := r.client.Incr(key).Result()
	if err != nil {
		log.Error("redis incr error---->", err)
		return 0, err
	}
	return result, nil
}

// 增加指定的数
func (r *RedisClient) IncrBy(key string, num int64) (int64, error) {
	result, err := r.client.IncrBy(key, num).Result()
	if err != nil {
		log.Error("redis IncrBy error ---->", err)
		return 0, err
	}
	return result, nil
}

// 执行递减操作
func (r *RedisClient) Decr(key string) (int64, error) {
	result, err := r.client.Decr(key).Result()
	if err != nil {
		log.Error("redis Decr error --->", err)
		return 0, err
	}
	return result, nil
}

// 减去指定的数
func (r *RedisClient) DecrBy(key string, num int64) (int64, error) {
	result, err := r.client.DecrBy(key, num).Result()
	if err != nil {
		log.Error("redis DecrBy error", result)
		return 0, err
	}
	r.client.Close()
	return result, nil
}

// 删除
func (r *RedisClient) Del(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		log.Error("redis  del error --->", err)
		return err
	}
	return nil
}

// 设置过期时间
func (r *RedisClient) Expire(key string, num time.Duration) error {
	err := r.client.Expire(key, num).Err()
	if err != nil {
		log.Error("redis Expire error--->", err)
		return err
	}
	return nil
}

// 查看key是否存在若 key 存在，返回 1 ，否则返回 0 。
func (r *RedisClient) Exists(key string) (int64, error) {
	result, err := r.client.Exists(key).Result()
	if err != nil {
		log.Error("redis exists error --->", err)
		return result, err
	}
	return result, nil

}
