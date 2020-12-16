package models

import (
	"time"
)

type SysDictData struct {
	DictCode   int64     `xorm:"not null pk autoincr comment('字典编码') BIGINT(20)"`
	DictSort   int       `xorm:"default 0 comment('字典排序') INT(4)"`
	DictLabel  string    `xorm:"default '' comment('字典标签') VARCHAR(100)"`
	DictValue  string    `xorm:"default '' comment('字典键值') VARCHAR(100)"`
	DictType   string    `xorm:"default '' comment('字典类型') VARCHAR(100)"`
	CssClass   string    `xorm:"comment('样式属性（其他样式扩展）') VARCHAR(100)"`
	ListClass  string    `xorm:"comment('表格回显样式') VARCHAR(100)"`
	IsDefault  string    `xorm:"default 'N' comment('是否默认（Y是 N否）') CHAR(1)"`
	Status     string    `xorm:"default '0' comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark     string    `xorm:"comment('备注') VARCHAR(500)"`
}
