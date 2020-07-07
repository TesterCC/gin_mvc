# 快速上手Go Project 1: Gin+Vue项目开发

适合后端开发快速上手Gin+Vue项目(Backend Part)。

## 课程地址：
[在线视频](https://www.bilibili.com/video/BV1CE411H7bQ)

## 安装包

```markdown
gin:
go get -u github.com/gin-gonic/gin

gorm:
go get -u github.com/jinzhu/gorm

go-sql-driver:
go get -u github.com/go-sql-driver/mysql

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

## Testing

1.Postman测试参数注意
- POST方法
- Body传参
- Body参数格式form-data

## Docs

- [gin](https://gin-gonic.com/docs/)
- [gorm](https://gorm.io/docs/)