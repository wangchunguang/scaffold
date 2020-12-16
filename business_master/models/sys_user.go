package models

import (
	"time"
)

type SysUser struct {
	UserId      int64     `xorm:"not null pk autoincr comment('用户ID') BIGINT(20)"`
	DeptId      int64     `xorm:"comment('部门ID') BIGINT(20)"`
	LoginName   string    `xorm:"not null comment('登录账号') VARCHAR(30)"`
	UserName    string    `xorm:"not null comment('用户昵称') VARCHAR(30)"`
	UserType    string    `xorm:"default '00' comment('用户类型（00系统用户）') VARCHAR(2)"`
	Email       string    `xorm:"default '' comment('用户邮箱') VARCHAR(50)"`
	Phonenumber string    `xorm:"default '' comment('手机号码') VARCHAR(11)"`
	Sex         string    `xorm:"default '0' comment('用户性别（0男 1女 2未知）') CHAR(1)"`
	Avatar      string    `xorm:"default '' comment('头像路径') VARCHAR(100)"`
	Password    string    `xorm:"default '' comment('密码') VARCHAR(50)"`
	Salt        string    `xorm:"default '' comment('盐加密') VARCHAR(20)"`
	Status      string    `xorm:"default '0' comment('帐号状态（0正常 1停用）') CHAR(1)"`
	DelFlag     string    `xorm:"default '0' comment('删除标志（0代表存在 2代表删除）') CHAR(1)"`
	LoginIp     string    `xorm:"default '' comment('最后登陆IP') VARCHAR(50)"`
	LoginDate   time.Time `xorm:"comment('最后登陆时间') DATETIME"`
	CreateBy    string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime  time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy    string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime  time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark      string    `xorm:"comment('备注') VARCHAR(500)"`
}
