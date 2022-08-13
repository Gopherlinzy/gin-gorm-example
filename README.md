# gin-gorm-example
这是一个简单的gin-gorm API项目，有很多有用的特性

## 安装
```
go get github.com/Gopherlinzy/gin-gorm-example
```

# 如何运行
***
## 需要
- Mysql
- go

## 准备
创建一个新的 blog 数据库,代码里有自动迁移数据表,会自动创建数据表

## 配置
修改 conf/app.ini 配置文件

```
[database]
Type = mysql
User = 数据库名称
Password = 数据库密码
//数据库的IP地址
Host = 127.0.0.1:3306
//数据库名称 
Name = blog
TablePrefix = blog_
```

## 运行
```
//进入到项目的里
cd $GOPATH/src/go-gin-example
//运行main文件
go run main.go 
```

显示项目运行信息和创建的API
```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/Gopherlinzy/go-gin-example/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/Gopherlinzy/go-gin-example/routers/api/v1.DeleteArticle (4 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (3 handlers)
```

# Swagger 文档
![image](https://user-images.githubusercontent.com/109471496/184467177-239dbd7c-0b3d-4a79-abff-bc8fd0dc678f.png)

# 特性
- RESTful API
- Gin
- jwt-go
- Gorm
- Swagger
- logging
- Graceful restart or stop (fvbock/endless)
- App configurable
