package main

import (
	"ginessential/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // 【import _ 包路径】只是引用该包，仅仅是为了调用init()函数，所以无法通过包名来调用包中的其他函数。
)



func main() {
	// IDE可以从这debug启动项目
	// 获取初始化后的db并延迟关闭
	db := common.InitDB()
	defer db.Close()

	r:=gin.Default()

	r = CollectRoute(r)
	panic(r.Run())  // listen and serve on 0.0.0.0:8080
}






