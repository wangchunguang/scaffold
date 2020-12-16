package redis

import log "github.com/sirupsen/logrus"

func (r *RedisClient) HSet(key, field, value string) error {
	err := r.client.HSet(key, field, value).Err()
	if err != nil {
		log.Error("redis HSet error--->", err)
		return err
	}
	return nil
}

func (r *RedisClient) HGet(key, field string) (string, error) {
	result, err := r.client.HGet(key, field).Result()
	if err != nil {
		log.Error("redis HGet error --->", err)
		return "", err
	}
	return result, nil
}

func (r *RedisClient) HGetAll(key string) (map[string]string, error) {
	result, err := r.client.HGetAll(key).Result()
	if err != nil {
		log.Error("redis HGetAll error --->", err)
		return nil, err
	}
	return result, nil
}

// 根据key和field字段，累加字段的数值
func (r *RedisClient) HIncrBy(key, field string, num int64) (int64, error) {
	result, err := r.client.HIncrBy(key, field, num).Result()
	if err != nil {
		log.Error("redis HIncrBy error--->", err)
		return 0, err
	}
	return result, nil
}

// 根据key获取所有的字段名
func (r *RedisClient) HKeys(key string) ([]string, error) {
	result, err := r.client.HKeys(key).Result()
	if err != nil {
		log.Error("redis HKeys error--->", err)
		return nil, err
	}
	return result, nil
}

func (r *RedisClient) HLen(key string) (int64, error) {
	result, err := r.client.HLen(key).Result()
	if err != nil {
		log.Error("redis HLen error--->", err)
		return 0, err
	}
	return result, nil
}

// 根据key和多个字段名，批量查询多个hash的值
func (r *RedisClient) HMGet(key string, fields ...string) ([]string, error) {
	result, err := r.client.HMGet(key, fields...).Result()
	if err != nil {
		log.Error("redis HMGet error --->", err)
		return nil, err
	}
	return InterfaceTOString(result), nil
}

// 根据key和多个字段名和字段值，批量设置hash字段值
func (r *RedisClient) HMSet(key string, data map[string]string) error {
	err := r.client.HMSet(key, data).Err()
	if err != nil {
		log.Error("redis HMSet error--->", err)
		return err
	}
	return nil
}

// 如果field字段不存在，则设置hash字段值
func (r *RedisClient) HSetNX(key, field string, num int64) error {
	err := r.client.HSetNX(key, field, num).Err()
	if err != nil {
		log.Error("redis HSetNX error --->", err)
		return err
	}
	return nil
}

func (r *RedisClient) HDel(key string, field ...string) (int64, error) {
	result, err := r.client.HDel(key, field...).Result()
	if err != nil {
		log.Error("redis HDel error --->", err)
		return 0, err
	}
	return result, nil
}

func (r *RedisClient) HExists(key, field string) (bool, error) {
	result, err := r.client.HExists(key, field).Result()
	if err != nil {
		log.Error("redis HExists error---> ", err)
		return result, err
	}
	return result, nil

}
