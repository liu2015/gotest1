package system

import (
	"ginserver/global"

	"strconv"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 权限相关的空结构体，作为对象。设置通用的方法

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityId := strconv.Itoa(int(AuthorityID))
	casbinService
}

func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success
}

// 使用方法清除数据
func (casbinService *CasbinService) RemoveFilteredPolicy(db *gorm.DB, authorityId string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0=?", authorityId).Error
}

// 同步数据库操作
func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, authorityId string, rules [][]string) error {
	err := casbinService.RemoveFilteredPolicy(db, authorityId)
	if err != nil {
		return err
	}
	return casbinService.AddPolicies(db, rules)

}

func (CasbinService *CasbinService) FreshCasbin() (err error) {
	e := CasbinService.Casbin()
	err = e.LoadPolicy()
	return err
}

// 清除匹配的权限
func (casbinService *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var CasbinRule []gormadapter.CasbinRule
	for v := range rules {
		CasbinRule = append(CasbinRule, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[v][0],
			V1:    rules[v][1],
			V2:    rules[v][2],
		})
	}
	return db.Create(&CasbinRule).Error

}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// 检查数据库否是innodb
func (CasbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {

	once.Do(func() {

		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败请检查casbin表是否是InnoDB引擎!", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败!", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()

	})
	return syncedCachedEnforcer
}
