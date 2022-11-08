<p align="center">
    GoRedisAdmin
</p>

## 简介

[go-redis-admin](https://github.com/linkaias/goRedisAdmin) 是一款使用golang和vue2开发的Redis的后台管理平台，它包含了在线的数据库管理和简洁的操作界面，目前项目处于开发阶段，后续功能会愈加完善。

- [在线预览](http://gradmin.uiucode.com/) 用户名：admin 密码：123456

- [使用文档](https://github.com/linkaias/goRedisAdmin) 文档开发中...

- 示例页面
<p align="center">
    <img width="900" src="http://gradmin.uiucode.com/image/home.png">
</p>
<p align="center">
    <img width="900" src="http://gradmin.uiucode.com/image/add.png">
</p>

- 备注
`项目开发中，非最终版本`

## 安装使用
项目处于开发版本，暂未打包，后期会打包各个系统可用包，如想本地预览，需要安装golang参考下面的"开发"

## 功能

```
- 登录 / 注销 (注销功能未完善)

- 数据库列表
- 数据库Key管理
- 新增Key(目前支持string、list、set、zset、hash)
- Key过期时间配置
- 删除某个Key
- 清空数据库(flushdb)
- 清空全部库（flushall）


```

## 开发

```bash
# 克隆项目
git clone https://github.com/linkaias/goRedisAdmin.git

# 进入项目目录
cd goRedisAdmin

# 根据注释修改配置文件
vim config.ini

# 安装依赖
go mod tidy

# 启动服务
go run main.go
```

浏览器访问 http://127.0.0.1:9527


## Online Demo

[在线 Demo](http://gradmin.uiucode.com/#/home)


## License

[MIT](https://github.com/linkaias/goRedisAdmin/blob/main/LICENSE)

Copyright (c) 2022-present LinKai
