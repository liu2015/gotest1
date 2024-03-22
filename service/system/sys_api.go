package system

import (
	"errors"
	"ginserver/global"
	"ginserver/model/system"

	"gorm.io/gorm"
)

type ApiService struct {
}

var ApiServiceApp = new(ApiService)

func (ApiService *ApiService) CreateApi(api system.SysApi) (err error) {

	if !errors.Is(global.GVA_DB.Where("path = ? AND method =?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.GVA_DB.Create(&api).Error
}

// 删除基础api
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GVA_DB.Where("id=?", api.ID).First(&entity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}

	// 鉴权操作
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil

}
