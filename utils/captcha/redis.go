package captcha

import (
	"context"
	"ginserver/global"
	"time"

	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type RedisStore struct {
	Expiration time.Duration
	Prekey     string
	Context    context.Context
}

func NewDefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		Prekey:     "CAPTCHA_",
		Context:    context.TODO(),
	}
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {
	err := global.GVA_REDIS.Set(rs.Context, rs.Prekey+id, value, rs.Expiration).Err()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreError!", zap.Error(err))
		return err

	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := global.GVA_REDIS.Get(rs.Context, key).Result()
	if err != nil {
		global.GVA_LOG.Error("RedisStoreGetError!", zap.Error(err))

		return ""
	}
	if clear {
		err := global.GVA_REDIS.Del(rs.Context, key).Err()
		if err != nil {
			global.GVA_LOG.Error("RedisStoreClear Error!", zap.Error(err))
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.Prekey + id
	v := rs.Get(key, clear)
	return v == answer
}
