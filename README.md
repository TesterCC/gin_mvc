# 快速上手Go Project 1: Gin+Vue项目开发

适合后端开发快速上手Gin+Vue项目(Backend Part)。

项目结构是基于MVC思想组织。

前端框架：Vue

后端框架：Gin

## 课程地址：

[参考视频](https://www.bilibili.com/video/BV1CE411H7bQ)

[目前进度](https://www.bilibili.com/video/BV1CE411H7bQ?p=3)

## 安装包

```markdown
gin:
go get -u github.com/gin-gonic/gin

gorm:
go get -u github.com/jinzhu/gorm

go-sql-driver:
go get -u github.com/go-sql-driver/mysql

jwt-go:
go get -u github.com/dgrijalva/jwt-go

config组件:
go get -u github.com/spf13/viper

go版本较低可能需要手动安装crypto：
go get -u golang.org/x/crypto

create mysql database:
mysql> show create database ginessential;
CREATE DATABASE `ginessential` /*!40100 DEFAULT CHARACTER SET utf8 */
```

## 关于编译
```
1.直接运行代码
cd ~/ginEssential
go run main.go routes.go


2.编译成二进制文件
go build
./ginessential
或者一起执行
go build && ./ginessential
```

## 关于token

token的组成以`.`为界分为3部分。
```markdown
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjYsImV4cCI6MTU5NDgxOTY1MCwiaWF0IjoxNTk0MjE0ODUwLCJpc3MiOiJmdWxsc3RhY2twZW50ZXN0LmNvbSIsInN1YiI6InVzZXIgdG9rZW4ifQ.x9lICXgHxvZ8WWKFnaLcnm9xzjB_sUMgUhJQGVJ8hS8

第1部分，协议头header，储存token使用的加密协议。
第2部分，载荷payload，存储&Claims中的信息。
第3部分，前2部分+jwtKey来hash的值

举例：
echo -n "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9" | base64 -D
{"alg":"HS256","typ":"JWT"}

echo -n "eyJVc2VySWQiOjYsImV4cCI6MTU5NDgxOTY1MCwiaWF0IjoxNTk0MjE0ODUwLCJpc3MiOiJmdWxsc3RhY2twZW50ZXN0LmNvbSIsInN1YiI6InVzZXIgdG9rZW4ifQ" | base64 -D
{"UserId":6,"exp":1594819650,"iat":1594214850,"iss":"fullstackpentest.com","sub":"user token"}
```

## Testing

接口测试使用postman即可

1.Postman测试参数注意
- POST方法
- Body传参
- Body参数格式form-data

服务端口默认设置为：1016，即 http://127.0.0.1:1016/api/xx/yy

测试注册用户接口范例：
![](https://cdn.jsdelivr.net/gh/TesterCC/pic_bed3/20220406225302.png)


## DB初始化

数据库驱动使用gorm

Docs: https://gorm.io/docs/

Install: `go get -u gorm.io/gorm`

### 创建数据库

1.图形化创建

New Database, 信息如下：

![](https://cdn.jsdelivr.net/gh/TesterCC/pic_bed3/20220406223812.png)

2.命令行创建

```mysql
CREATE DATABASE
IF
	NOT EXISTS ginessential DEFAULT CHARACTER 
	SET utf8 DEFAULT COLLATE utf8_general_ci;
```
实际生产环境：utf8mb4和utf8mb4_general_ci 用的更多。

PS：基本上mysql库和表创建成功就可以启动server服务器了。

## Docs
- [gin](https://gin-gonic.com/docs/)
- [gorm](https://gorm.io/docs/)
- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [Installing multiple Go versions](https://go.dev/doc/manage-install#installing-multiple)

