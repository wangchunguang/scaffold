package models

import (
	"time"
)

type SysJob struct {
	JobId          int64     `xorm:"not null pk autoincr comment('任务ID') BIGINT(20)"`
	JobName        string    `xorm:"not null pk default '' comment('任务名称') VARCHAR(64)"`
	JobGroup       string    `xorm:"not null pk default 'DEFAULT' comment('任务组名') VARCHAR(64)"`
	InvokeTarget   string    `xorm:"not null comment('调用目标字符串') VARCHAR(500)"`
	CronExpression string    `xorm:"default '' comment('cron执行表达式') VARCHAR(255)"`
	MisfirePolicy  string    `xorm:"default '3' comment('计划执行错误策略（1立即执行 2执行一次 3放弃执行）') VARCHAR(20)"`
	Concurrent     string    `xorm:"default '1' comment('是否并发执行（0允许 1禁止）') CHAR(1)"`
	Status         string    `xorm:"default '0' comment('状态（0正常 1暂停）') CHAR(1)"`
	CreateBy       string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime     time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy       string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime     time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark         string    `xorm:"default '' comment('备注信息') VARCHAR(500)"`
}
