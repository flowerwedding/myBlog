# myblog

[swagger接口文档](http://127.0.0.1:8000/swagger/index.html)

## 初衷

该项目模仿《go 语言编程之旅》的博客后台，并修改和增加了部分功能。

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
    |-docs 文档目录
    |-global 全局变量目录
    |-internal 内部模块目录
        |-dao 数据访问目录
        |-middleware HTTP中间件目录
        |-model 模型目录
        |-routers 路由相关逻辑目录
            |-api 接口目录
        |-service 项目核心业务逻辑目录
    |-pkg 项目相关模块包目录
        |-app 基本功能
        |-captcha 图片验证码
        |-convert 类型转换
        |-email 邮件
        |-errcode 错误码
        |-limiter 限流器
        |-logger 日志
        |-setting 配置
        |-tracer 链路追踪
        |-upload 上传文件
        |-util 其他
    |-scripts 各类构建、安装、分析等操作的脚本目录
    |-storage 项目生成的临时文件目录
        |-logs 项目日志目录
        |-uploads 项目上传的文件目录
    |-third_party 第三方资源工具目录
    |-vendor 项目依赖其他开源项目目录
    |-view 模板文件目录
    |-main.go 程序执行入口
    |-setup.go
```

## TODO

- [ ] 文章、页面访问统计
- [ ] GitHub登录发表评论
- [x] rss订阅
- [x] 图形验证码
- [x] ~~七牛云~~SMMS图床
- [x] 邮箱提醒功能
- [x] 链路追踪
- [x] 统一超时、接口限流控制
- [x] 系统日志
- [x] 优雅重启和停止
- [ ] Markdown

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
3. 如果需要使用邮件功能，请自行填写
   - smtp_username
   - smtp_password
   - smtp_from
   - smtp_to

## 效果图

![1598147827957](https://github.com/flowerwedding/myBlog/blob/master/view/img/1598147827957.png)
