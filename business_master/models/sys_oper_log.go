package models

import (
	"time"
)

type SysOperLog struct {
	OperId        int64     `xorm:"not null pk autoincr comment('日志主键') BIGINT(20)"`
	Title         string    `xorm:"default '' comment('模块标题') VARCHAR(50)"`
	BusinessType  int       `xorm:"default 0 comment('业务类型（0其它 1新增 2修改 3删除）') INT(2)"`
	Method        string    `xorm:"default '' comment('方法名称') VARCHAR(100)"`
	RequestMethod string    `xorm:"default '' comment('请求方式') VARCHAR(10)"`
	OperatorType  int       `xorm:"default 0 comment('操作类别（0其它 1后台用户 2手机端用户）') INT(1)"`
	OperName      string    `xorm:"default '' comment('操作人员') VARCHAR(50)"`
	DeptName      string    `xorm:"default '' comment('部门名称') VARCHAR(50)"`
	OperUrl       string    `xorm:"default '' comment('请求URL') VARCHAR(255)"`
	OperIp        string    `xorm:"default '' comment('主机地址') VARCHAR(50)"`
	OperLocation  string    `xorm:"default '' comment('操作地点') VARCHAR(255)"`
	OperParam     string    `xorm:"default '' comment('请求参数') VARCHAR(2000)"`
	JsonResult    string    `xorm:"default '' comment('返回参数') VARCHAR(2000)"`
	Status        int       `xorm:"default 0 comment('操作状态（0正常 1异常）') INT(1)"`
	ErrorMsg      string    `xorm:"default '' comment('错误消息') VARCHAR(2000)"`
	OperTime      time.Time `xorm:"comment('操作时间') DATETIME"`
}
