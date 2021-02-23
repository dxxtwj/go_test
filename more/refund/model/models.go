package model

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 初始化数据库配置
func InitDb() *gorm.DB {
	var (
		err                                            error
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
	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName,
	))
	if err != nil {
		panic(err)
	}

	db.SingularTable(true)          //如果使用gorm来帮忙创建表时，这里填写false的话gorm会给表添加s后缀，填写true则不会

	//db.LogMode(true)                //打印sql语句

	//开启连接池
	db.DB().SetMaxIdleConns(100)        //最大空闲连接
	db.DB().SetMaxOpenConns(10000)      //最大连接数
	db.DB().SetConnMaxLifetime(30)      //最大生存时间(s)

	return db
}
