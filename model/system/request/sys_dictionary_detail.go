package request

import (
	"ginserver/model/common/request"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
