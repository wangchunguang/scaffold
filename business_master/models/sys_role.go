package models

import (
	"time"
)

type SysRole struct {
	RoleId     int64     `xorm:"not null pk autoincr comment('角色ID') BIGINT(20)"`
	RoleName   string    `xorm:"not null comment('角色名称') VARCHAR(30)"`
	RoleKey    string    `xorm:"not null comment('角色权限字符串') VARCHAR(100)"`
	RoleSort   int       `xorm:"not null comment('显示顺序') INT(4)"`
	DataScope  string    `xorm:"default '1' comment('数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）') CHAR(1)"`
	Status     string    `xorm:"not null comment('角色状态（0正常 1停用）') CHAR(1)"`
	DelFlag    string    `xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	CreateBy   string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy   string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark     string    `xorm:"comment('备注') VARCHAR(500)"`
}
