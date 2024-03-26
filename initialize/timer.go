package initialize

import (
	"fmt"
	"ginserver/global"

	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/robfig/cron/v3"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB)
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}
	}()
}
