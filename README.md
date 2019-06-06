# goshare

基于go标准库实现的 博客后端API服务。

[![Build Status](https://travis-ci.org/ligenhw/goshare.svg?branch=master)](https://travis-ci.org/ligenhw/goshare)
[![codecov](https://codecov.io/gh/ligenhw/goshare/branch/master/graph/badge.svg)](https://codecov.io/gh/ligenhw/goshare)
[![Go Report Card](https://goreportcard.com/badge/github.com/ligenhw/goshare)](https://goreportcard.com/report/github.com/ligenhw/goshare)
[![codebeat badge](https://codebeat.co/badges/ea8dd5a0-964f-4f34-8cae-c870629da46d)](https://codebeat.co/projects/github-com-ligenhw-goshare-master)

前端项目: [https://github.com/ligenhw/goshare-website](https://github.com/ligenhw/goshare-website)

## 安装

go get -u github.com/ligenhw/goshare

## 功能

* 文章
* 用户
* 评论
* 三方登录 github qq 支付宝

## 计划加入的功能

* 搜索
* 博客迁移
  支持迁移 简书,CSDN,博客园中的文章及其评论
  
## ⚙️ 配置 & 环境变量

config.json

export DSN="gen:1234@tcp(192.168.199.230)/goshare?charset=utf8mb4&parseTime=true"

export ADDRESS=":8080"

## 构建执行

GOOS=linux GOARCH=amd64 go build

./goshare

## Docker方式部署

* 1.构建镜像

docker build -t goshare .

* 2.启动容器

### mysql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=123 -d mysql

### goshare
docker run -d --name some-goshare --link some-mysql:db -e DSN="root:123@tcp(db)/goshare?charset=utf8&parseTime=true" goshare

> 三方登录的api secret需要换成正式的

### nginx
docker run --name some-nginx -p 80:80 -d -v  ~/goshare-website/build:/usr/share/nginx/html nginx

## Docker Compose 方式部署

cd contrib/compose

docker-compose up -d

---

## 改进点
* 使用context传递请求上下文参数，解除session, auth 与业务的耦合
>参考 https://www.ddhigh.com/2018/10/17/golang-context-with-value.html
