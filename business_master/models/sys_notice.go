package models

import (
	"time"
)

type SysNotice struct {
	NoticeId      int       `xorm:"not null pk autoincr comment('公告ID') INT(4)"`
	NoticeTitle   string    `xorm:"not null comment('公告标题') VARCHAR(50)"`
	NoticeType    string    `xorm:"not null comment('公告类型（1通知 2公告）') CHAR(1)"`
	NoticeContent string    `xorm:"comment('公告内容') VARCHAR(2000)"`
	Status        string    `xorm:"default '0' comment('公告状态（0正常 1关闭）') CHAR(1)"`
	CreateBy      string    `xorm:"default '' comment('创建者') VARCHAR(64)"`
	CreateTime    time.Time `xorm:"comment('创建时间') DATETIME"`
	UpdateBy      string    `xorm:"default '' comment('更新者') VARCHAR(64)"`
	UpdateTime    time.Time `xorm:"comment('更新时间') DATETIME"`
	Remark        string    `xorm:"comment('备注') VARCHAR(255)"`
}
