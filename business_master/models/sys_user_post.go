package models

type SysUserPost struct {
	UserId int64 `xorm:"not null pk comment('用户ID') BIGINT(20)"`
	PostId int64 `xorm:"not null pk comment('岗位ID') BIGINT(20)"`
}
