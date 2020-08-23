# myblog

[示例地址](http://67.216.221.42/)

## 初衷

该项目模仿《go 语言编程之旅》的博客后台，并增加/修改自己的理解。

还是想去红岩网校吧，真的好想。

## 技术选型

1. web:[gin](https://github.com/gin-gonic/gin)
2. orm:[gorm](https://github.com/jinzhu/gorm)
3. database:[MySQL](https://github.com/mattn/go-sqlite3)
4. ~~全文检索:[wukong](https://github.com/huichen/wukong)~~
5. 文件存储:~~[七牛云存储](https://www.qiniu.com/)~~[smms图床](https://sm.ms/)
6. 配置文件 [go-yaml](https://github.com/go-yaml/yaml)

## 项目结构

```
-myBlog
    |-configs 配置文件目录
    |-docs
    |-global
    |-internal
        |-dao
        |-middleware
        |-model
        |-routers
            |-api
        |-service
    |-pkg
        |-app
        |-convert
        |-email
        |-errcode
        |-limiter
        |-logger
        |-setting
        |-tracer
        |-upload
        |-util
    |-scripts
    |-storage
        |-logs
        |-uploads
    |-third_party
    |-vendor 项目依赖其他开源项目目录
    |-view
    |-main.go 程序执行入口
    |-setup.go
```

## TODO



## 安装部署

本项目使用govendor管理依赖包，[govendor](https://github.com/kardianos/govendor)安装方法

```
go get -u github.com/kardianos/govendor
```

```
git clone https://github.com/flowerwedding/myblog
cd myblog
govendor sync
go run main.go
```

## 使用方法

### 使用说明

使用 /auth 接口登录，每次发送请求携带所生成的 token，或放在请求头中

### 注意事项

1. 如果需求上传图片功能请自行申请七牛云存储空间，并修改配置文件填写
   -  qiniu_accesskey
   -  qiniu_secretkey
   -  qiniu_fileserver 七牛访问地址
   -  qiniu_bucket 空间名称
2. 如果需要github登录评论功能请自行注册github oauthapp，并修改配置文件填写
   - github_clientid
   - github_clientsecret
   - github_redirecturl
3. 如果需要使用邮件订阅功能，请自行填写
   - smtp_username
   - smtp_password
   - smtp_host,例如：smtp.163.com:25

## 效果图

![1598147827957](C:\Users\HUAWEI\myBlog\view\img\main.go)