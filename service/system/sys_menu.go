package system

import system "ginserver/model/response"

// 获得路由的总树map

type MenuService struct {
}

var MenuServiceApp = new(MenuService)

func (menuService *MenuService) getMenuTreeMap(authorityId uint) (treeMap map[string][]system.SysMenu, err error) {

}
