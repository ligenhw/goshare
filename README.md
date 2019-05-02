# goshare

开源博客系统

[![Build Status](https://travis-ci.org/ligenhw/goshare.svg?branch=master)](https://travis-ci.org/ligenhw/goshare)
[![codecov](https://codecov.io/gh/ligenhw/goshare/branch/master/graph/badge.svg)](https://codecov.io/gh/ligenhw/goshare)
[![Go Report Card](https://goreportcard.com/badge/github.com/ligenhw/goshare)](https://goreportcard.com/report/github.com/ligenhw/goshare)


## 安装

go get -u github.com/ligenhw/goshare

## 功能

* 文章
* 用户
* 评论
* 三方登录 github

## 计划加入的功能

* 三方登陆
  支持 qq , 微信
* 搜索
* 博客迁移
  支持迁移 简书,CSDN,博客园中的文章及其评论
* 博客爬虫
  定期从 简书,CSDN,博客园 获取热门文章
  
## ⚙️ 配置

config.json

## 环境变量

export DSN="gen:1234@tcp(192.168.199.231)/goshare?charset=utf8&parseTime=true"
export ADDRESS=":8080"

## 构建执行

GOOS=linux GOARCH=amd64 go build 
./goshare

## Docker

构建镜像
docker build -t goshare .

启动容器

### mysql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=123 -d mysql

### web
docker run -d --name some-goshare --link some-mysql:db -e DSN="root:123@tcp(db)/goshare?charset=utf8&parseTime=true" goshare

### nginx



## 改进点
* 使用context传递请求上下文参数，解除session, auth 与业务的耦合
>参考 https://www.ddhigh.com/2018/10/17/golang-context-with-value.html

