package controller

import (
	"ginessential/common"
	"ginessential/model"
	"ginessential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(ctx *gin.Context) {
	//引入DB实例
	DB := common.GetDB()

	// get args
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 数据验证
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"手机号必须为11位"})    // 422 Unprocessable Entity 请求格式正确,但是由于含有语义错误,无法响应。
		return
	}

	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"密码不能少于6位"})
		return
	}

	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name,telephone,password)
	// 判断手机号是否存在。 如果用户存在就不允许注册
	if isTelephoneExist(DB, telephone){
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code":422, "msg":"用户已存在"})
		return
	}

	// 创建用户
	newUser := model.User {
		Name: name,
		Telephone: telephone,
		Password: password,
	}
	DB.Create(&newUser)

	// 返回结果

	ctx.JSON(200, gin.H{
		"message": "注册成功",
	})
}




func isTelephoneExist(db *gorm.DB, telephone string) bool{
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}