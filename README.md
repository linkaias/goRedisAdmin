<p align="center">
    GoRedisAdmin
</p>

## 简介

[go-redis-admin](https://github.com/linkaias/goRedisAdmin)
是一款使用golang和vue2开发的Redis的后台管理平台，它包含了在线的数据库管理和简洁的操作界面，旨在方便用户管理redis数据。

- [在线预览](http://gradmin.uiucode.com/#/login) 用户名：admin 密码：123456

- [使用文档](https://github.com/linkaias/goRedisAdmin)

- 示例页面

<p align="center">
    <img width="900" src="http://gradmin.uiucode.com/image/login.png">
</p>
<p align="center">
    <img width="900" src="http://gradmin.uiucode.com/image/home.png">
</p>
<p align="center">
    <img width="900" src="http://gradmin.uiucode.com/image/add.png">
</p>

## 功能

```
- 登录/注销 
- 数据库列表
- 数据库Key管理
- 新增Key(目前支持string、list、set、zset、hash)
- Key过期时间配置
- 支持添加访问IP白名单 
- 删除某个Key
- 清空数据库(flushdb)
- 清空全部库（flushall）

```

## 安装

```bash
# 克隆项目
git clone https://github.com/linkaias/goRedisAdmin.git

# 进入项目目录
cd goRedisAdmin

# 创建日志文件夹,此文件夹需要可写权限
mkdir var

# 根据注释修改配置文件
vim config.ini

# 安装依赖
go mod tidy

# 启动服务
go run main.go
```

部署完成后本地浏览器访问 http://127.0.0.1:9527

## Online Demo

[在线 Demo](http://gradmin.uiucode.com/#/login)
[Online Demo](http://gradmin.uiucode.com/#/login)

## JetBrains open source certificate support

The GoRedisAdmin project has always been developed in the GoLand integrated development environment under JetBrains,
based on the free JetBrains Open Source license(s) genuine free license. I would like to express my gratitude.

<a href="https://www.jetbrains.com/">
<img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" />
</a>

## License

[MIT](https://github.com/linkaias/goRedisAdmin/blob/main/LICENSE)

Copyright (c) 2022-present LinKai
