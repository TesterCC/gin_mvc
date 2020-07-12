package controller

import (
	"ginessential/common"
	"ginessential/dto"
	"ginessential/model"
	"ginessential/response"
	"ginessential/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"}) // 422 Unprocessable Entity 请求格式正确,但是由于含有语义错误,无法响应。
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否存在。 如果用户存在就不允许注册
	if isTelephoneExist(DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已存在")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}

	// 创建用户
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		//ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密错误"})
		return
	}

	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashedPassword), // hashed password
	}
	DB.Create(&newUser)

	// 返回结果

	//ctx.JSON(http.StatusOK, gin.H{
	//	"code":    200,
	//	"message": "注册成功",
	//})

	// 重新封装
	response.Success(ctx, nil, "注册成功")
}

func Login(ctx *gin.Context) {
	DB := common.GetDB()
	// 获取参数
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	log.Println(telephone, password)
	// 数据验证
	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"}) // 422 Unprocessable Entity 请求格式正确,但是由于含有语义错误,无法响应。
		return
	}

	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码不能少于6位"})
		return
	}

	// 判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		// ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
	}

	// 判断密码是否正确   用户密码不能明文保存，所以在注册时应该将用户密码加密，使用密文存储
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Fail(ctx, nil, "密码错误")
		//ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
	}

	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error : %v ", err)
		return
	}

	// 返回结果
	response.Success(ctx, gin.H{"token":token},"登录成功")
	//ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "登录成功", "data": gin.H{"token": token}})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}


func Info(ctx *gin.Context) {
	// 获取用户信息时，用户应该是通过认证的，所以应该能够直接从上下文中提取用户信息
	user, exist := ctx.Get("user")

	// 自己加的，判断万一出现的没有获取到的情况
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "上下文信息获取失败"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"code":200, "data":gin.H{"user":dto.ToUserDto(user.(model.User))}})   // 将user转成user dto
	}

}