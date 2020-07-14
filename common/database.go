package common

import (
	"fmt"
	"ginessential/model"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)


var DB *gorm.DB

// 数据库初始化
func InitDB() *gorm.DB {
	//drivername := "mysql"
	//host := "localhost"
	//port := "3306"
	//database := "ginessential"
	//username := "root"
	//password := ""
	//charset := "utf8"

	//使用viper获取config info，注意key写法参考application.yml
	drivername := viper.GetString("datasource.driverName")
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")

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