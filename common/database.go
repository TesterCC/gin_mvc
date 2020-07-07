package common

import (
	"fmt"
	"ginessential/model"
	"github.com/jinzhu/gorm"
	"log"
)


var DB *gorm.DB

// 数据库初始化
func InitDB() *gorm.DB {
	drivername := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginessential"
	username := "root"
	password := "yanxi76543210"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	log.Println(args)

	db, err := gorm.Open(drivername, args)
	if err!=nil {
		panic("failed to connect database, err: " + err.Error())
	}

	// 让gorm自动创建数据表
	db.AutoMigrate(&model.User{})

	DB = db
	return db
}

//返回DB实例
func GetDB() *gorm.DB {
	return DB
}