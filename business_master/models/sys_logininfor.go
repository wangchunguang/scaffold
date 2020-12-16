package models

import (
	"time"
)

type SysLogininfor struct {
	InfoId        int64     `xorm:"not null pk autoincr comment('访问ID') BIGINT(20)"`
	LoginName     string    `xorm:"default '' comment('登录账号') VARCHAR(50)"`
	Ipaddr        string    `xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation string    `xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser       string    `xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os            string    `xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status        string    `xorm:"default '0' comment('登录状态（0成功 1失败）') CHAR(1)"`
	Msg           string    `xorm:"default '' comment('提示消息') VARCHAR(255)"`
	LoginTime     time.Time `xorm:"comment('访问时间') DATETIME"`
}
