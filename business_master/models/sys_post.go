package models

import (
	"time"
)

type SysPost struct {
	PostId     int64     `xorm:"not null pk autoincr comment('岗位ID') BIGINT(20)"`
	PostCode   string    `xorm:"not null comment('岗位编码') VARCHAR(64)"`
	PostName   string    `xorm:"not null comment('岗位名称') VARCHAR(50)"`
	PostSort   int       `xorm:"not null comment('显示顺序') INT(4)"`
	Status     string    `xorm:"not null comment('状态（0正常 1停用）') CHAR(1)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark     string    `xorm:"comment('备注') VARCHAR(500)"`
}
