package main

import (
	"fmt"
	"ginessential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。
	"github.com/spf13/viper"
	"os"
)

func main() {
	//在项目一开始就读取配置
	InitConfig()

	// IDE可以从这debug启动项目
	// 获取初始化后的db并延迟关闭
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	r = CollectRoute(r)

	// 利用配置文件yml，改监听端口，key写法参考application.yml
	port := viper.GetString("server.port")

	if port != "" {
		fmt.Printf("Start server at port %v\n", port)
		panic(r.Run(":" + port))
	}

	panic(r.Run()) // default listen and serve on 0.0.0.0:8080
}

func InitConfig() {
	// 定义函数获取当前的工作目录
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")

	errinfo := viper.ReadInConfig()
	if errinfo != nil {
		panic(errinfo)
	}

}
