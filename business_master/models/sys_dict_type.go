package models

import (
	"time"
)

type SysDictType struct {
	DictId     int64     `xorm:"not null pk autoincr comment('字典主键') BIGINT(20)"`
	DictName   string    `xorm:"default '' comment('字典名称') VARCHAR(100)"`
	DictType   string    `xorm:"default '' comment('字典类型') unique VARCHAR(100)"`
	Status     string    `xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark     string    `xorm:"comment('备注') VARCHAR(500)"`
}
