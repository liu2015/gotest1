package request

import (
	"ginserver/model/common/request"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
