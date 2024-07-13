package captcha

import (
	"context"
	"go.uber.org/zap"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/global"
	"time"
)

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: 180 * time.Second,
		PreKey:     "Captcha_",
		Context:    context.TODO(),
	}
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.App.RedisClient.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	if err != nil {
		logger.Log.Error("RedisStoreSetError: ", zap.Error(err))
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.App.RedisClient.Get(rs.Context, key).Result()
	if err != nil {
		logger.Log.Error("RedisStoreGetError:", zap.Error(err))
		return ""
	}
	if clear {
		err := global.App.RedisClient.Del(rs.Context, key).Err()
		if err != nil {
			logger.Log.Error("RedisStoreClearError!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	v := rs.Get(key, clear)
	return v == answer
}
