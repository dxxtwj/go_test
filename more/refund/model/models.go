package model

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID int64 `gorm:"primary_key" json:"id"`
}

// 配置成功后的db
var db *gorm.DB

// 初始化数据库配置
func init() {
	var (
		err error
		typeName, user, password, host, dbName, dbType string
	)
	cfg, err := ini.Load("config/env.ini")
	if err != nil {
		panic(err)
	}
	typeName = "database"
	user = cfg.Section(typeName).Key("USER").String()
	password = cfg.Section(typeName).Key("PASSWORD").String()
	host = cfg.Section(typeName).Key("HOST").String()
	dbName = cfg.Section(typeName).Key("NAME").String()
	dbType = cfg.Section(typeName).Key("TYPE").String()

	// db弄成全局变量
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
		))
	if err != nil {
		panic(err)
	}
	defer db.Close()
}