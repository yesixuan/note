package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"note/src/util"
	"strings"
	"time"
)

var DB *gorm.DB

func init() {
	path := strings.Join([]string{util.Configs.Mysql.Name, ":", util.Configs.Mysql.Pwd, "@(", util.Configs.Mysql.Ip, ":", util.Configs.Mysql.Port, ")/", util.Configs.Mysql.DBName, "?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	db, err := gorm.Open("mysql", path)
	//defer DB.Close()
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.DB().SetConnMaxLifetime(1 * time.Second)
	db.DB().SetMaxIdleConns(20)   //最大打开的连接数
	db.DB().SetMaxOpenConns(2000) //设置最大闲置个数
	db.SingularTable(true)        //表生成结尾不带s
	// 启用Logger，显示详细日志
	db.LogMode(true)
	if !db.HasTable(&User{}) { //db.Set 设置一些额外的表属性                              //db.CreateTable创建表
		if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").Error; err != nil {
			panic(err)
		}
	}
	DB = db
}

func InitTables() {
	// 自动创建表
	DB.AutoMigrate(
		&User{},
	)
}
