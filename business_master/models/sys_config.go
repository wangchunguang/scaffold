package models

import (
	"time"
)

type SysConfig struct {
	ConfigId    int       `xorm:"not null pk autoincr comment('参数主键') INT(5)"`
	ConfigName  string    `xorm:"default '' comment('参数名称') VARCHAR(100)"`
	ConfigKey   string    `xorm:"default '' comment('参数键名') VARCHAR(100)"`
	ConfigValue string    `xorm:"default '' comment('参数键值') VARCHAR(500)"`
	ConfigType  string    `xorm:"default 'N' comment('系统内置（Y是 N否）') CHAR(1)"`
	CreateBy    string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark      string    `xorm:"comment('备注') VARCHAR(500)"`
}
