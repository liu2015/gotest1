package system

import (
	"context"
	"ginserver/global"
	"ginserver/model/system"
	"ginserver/utils"

	"go.uber.org/zap"
)

type JwtService struct {
}

func (JwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {

	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

func (JwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// 获得redis 鉴权
func (jwtService *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

func (JwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}

	timer := dr

	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error

	if err != nil {
		global.GVA_LOG.Error("加载数据库jwt黑名单失败！", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
