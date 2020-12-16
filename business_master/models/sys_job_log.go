package models

import (
	"time"
)

type SysJobLog struct {
	JobLogId      int64     `xorm:"not null pk autoincr comment('任务日志ID') BIGINT(20)"`
	JobName       string    `xorm:"not null comment('任务名称') VARCHAR(64)"`
	JobGroup      string    `xorm:"not null comment('任务组名') VARCHAR(64)"`
	InvokeTarget  string    `xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	JobMessage    string    `xorm:"comment('日志信息') VARCHAR(500)"`
	Status        string    `xorm:"default '0' comment('执行状态（0正常 1失败）') CHAR(1)"`
	ExceptionInfo string    `xorm:"default '' comment('异常信息') VARCHAR(2000)"`
	CreateTime    time.Time `xorm:"comment('创建时间') DATETIME"`
}
