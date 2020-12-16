package models

import (
	"time"
)

type SysMenu struct {
	MenuId     int64     `xorm:"not null pk autoincr comment('菜单ID') BIGINT(20)"`
	MenuName   string    `xorm:"not null comment('菜单名称') VARCHAR(50)"`
	ParentId   int64     `xorm:"default 0 comment('父菜单ID') BIGINT(20)"`
	OrderNum   int       `xorm:"default 0 comment('显示顺序') INT(4)"`
	Url        string    `xorm:"default '#' comment('请求地址') VARCHAR(200)"`
	Target     string    `xorm:"default '' comment('打开方式（menuItem页签 menuBlank新窗口）') VARCHAR(20)"`
	MenuType   string    `xorm:"default '' comment('菜单类型（M目录 C菜单 F按钮）') CHAR(1)"`
	Visible    string    `xorm:"default '0' comment('菜单状态（0显示 1隐藏）') CHAR(1)"`
	Perms      string    `xorm:"comment('权限标识') VARCHAR(100)"`
	Icon       string    `xorm:"default '#' comment('菜单图标') VARCHAR(100)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark     string    `xorm:"default '' comment('备注') VARCHAR(500)"`
}
