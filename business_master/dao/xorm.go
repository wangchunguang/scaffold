package dao

import (
	"business_master/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	log "github.com/sirupsen/logrus"
)

var engine *xorm.Engine

func MysqlInit(yamlConfig *config.YamlConfig) error {
	engine, err := xorm.NewEngine("mysql", yamlConfig.Mysql.User+":"+yamlConfig.Mysql.Password+"@("+
		yamlConfig.Mysql.Host+":"+yamlConfig.Mysql.Port+")/"+yamlConfig.Mysql.Name+"?charset=utf8")
	if err != nil {
		log.Error(err)
		return err
	}
	err = engine.Ping()
	if err != nil {
		log.Error("mysql ping error --->", err)
		return err
	}
	// sql输出在控制台
	engine.ShowSQL(true)
	// 空闲连接数
	engine.SetMaxIdleConns(3)
	// 最大连接数
	engine.SetMaxOpenConns(100)
	//engine.Logger().SetLevel(core.LOG_ERR)
	return nil
}
