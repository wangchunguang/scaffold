package models

import (
	"time"
)

type SysDept struct {
	DeptId     int64     `xorm:"not null pk autoincr comment('部门id') BIGINT(20)"`
	ParentId   int64     `xorm:"default 0 comment('父部门id') BIGINT(20)"`
	Ancestors  string    `xorm:"default '' comment('祖级列表') VARCHAR(50)"`
	DeptName   string    `xorm:"default '' comment('部门名称') VARCHAR(30)"`
	OrderNum   int       `xorm:"default 0 comment('显示顺序') INT(4)"`
	Leader     string    `xorm:"comment('负责人') VARCHAR(20)"`
	Phone      string    `xorm:"comment('联系电话') VARCHAR(11)"`
	Email      string    `xorm:"comment('邮箱') VARCHAR(50)"`
	Status     string    `xorm:"default '0' comment('部门状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
}
