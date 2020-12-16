package models

import (
	"time"
)

type SysUserOnline struct {
	Sessionid      string    `xorm:"not null pk default '' comment('用户会话id') VARCHAR(50)"`
	LoginName      string    `xorm:"default '' comment('登录账号') VARCHAR(50)"`
	DeptName       string    `xorm:"default '' comment('部门名称') VARCHAR(50)"`
	Ipaddr         string    `xorm:"default '' comment('登录IP地址') VARCHAR(50)"`
	LoginLocation  string    `xorm:"default '' comment('登录地点') VARCHAR(255)"`
	Browser        string    `xorm:"default '' comment('浏览器类型') VARCHAR(50)"`
	Os             string    `xorm:"default '' comment('操作系统') VARCHAR(50)"`
	Status         string    `xorm:"default '' comment('在线状态on_line在线off_line离线') VARCHAR(10)"`
	StartTimestamp time.Time `xorm:"comment('session创建时间') DATETIME"`
	LastAccessTime time.Time `xorm:"comment('session最后访问时间') DATETIME"`
	ExpireTime     int       `xorm:"default 0 comment('超时时间，单位为分钟') INT(5)"`
}
