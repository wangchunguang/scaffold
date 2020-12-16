package models

type SysUserRole struct {
	UserId int64 `xorm:"not null pk comment('用户ID') BIGINT(20)"`
	RoleId int64 `xorm:"not null pk comment('角色ID') BIGINT(20)"`
}
