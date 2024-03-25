package router

import (
	"ginserver/service/example"
	"ginserver/service/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}
