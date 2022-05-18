<h1 align="center"><a href="https://www.bestlang.cn" target="_blank">goshare</a></h1>

> 基于go标准库实现的 博客后端API服务

[![GitHub Actions CI][ciBadge]][ciLink]
[![codecov](https://codecov.io/gh/ligenhw/goshare/branch/master/graph/badge.svg)](https://codecov.io/gh/ligenhw/goshare)
[![Go Report Card](https://goreportcard.com/badge/github.com/ligenhw/goshare)](https://goreportcard.com/report/github.com/ligenhw/goshare)
[![codebeat badge](https://codebeat.co/badges/ea8dd5a0-964f-4f34-8cae-c870629da46d)](https://codebeat.co/projects/github-com-ligenhw-goshare-master)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fligenhw%2Fgoshare.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fligenhw%2Fgoshare?ref=badge_shield)

前端项目: [https://github.com/ligenhw/goshare-website](https://github.com/ligenhw/goshare-website)

[ciBadge]: https://github.com/ligenhw/goshare/actions/workflows/go.yml/badge.svg
[ciLink]: https://github.com/ligenhw/goshare/actions/workflows/go.yml

## Contents 目录

- [Introduction 介绍 ✨](#introduction-介绍-)
- [功能 🔥](#功能-)
- [计划加入的功能 🎉](#计划加入的功能-)
- [配置 & 环境变量️️ ️⚙️](#配置--环境变量️️-️️)
- [构建执行 📦](#构建执行-)
- [Docker方式部署 ✈️](#docker方式部署-️)
- [Docker Compose 方式部署 🚀](#docker-compose-方式部署-)
- [Show your support ⭐️](#Show-your-support-)
- [License 📝](#License-)

## Introduction 介绍 ✨

goshare is a blog web api service by golang.

goshare 是基于go标准库实现的 博客后端API服务。


## 功能 🔥

* 文章
* 用户
* 评论
* 标签
* 三方登录 github qq 支付宝

## 计划加入的功能 🎉

* 搜索
* 博客迁移
  支持迁移 简书,CSDN,博客园中的文章及其评论


## 配置 & 环境变量️️ ️⚙️

生产环境: configration/config.json  
开发环境: configration/config.dev.json 

环境变量:

```bash
export DSN="gen:1234@tcp(192.168.199.230)/goshare?charset=utf8mb4&parseTime=true"
export ADDRESS=":8080"
```

## 构建执行 📦

```bash
GOOS=linux GOARCH=amd64 go build

./goshare
```

## Docker方式部署 ✈️

* 1.构建镜像
```bash
docker build -t goshare .
```

* 2.启动容器

### mysql
```bash
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=123 -d mysql
```

### goshare
```bash
docker run -d --name some-goshare --link some-mysql:db -e DSN="root:123@tcp(db)/goshare?charset=utf8&parseTime=true" goshare
```

> 三方登录的api secret需要换成正式的

### nginx
```bash
docker run --name some-nginx -p 80:80 -d -v  ~/goshare-website/build:/usr/share/nginx/html nginx
```

## Docker Compose 方式部署 🚀

```bash
cd contrib/compose

docker-compose up -d
```

## Show your support ⭐️

Please ⭐️ this repository if this project helped you!

---

## License 📝
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fligenhw%2Fgoshare.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fligenhw%2Fgoshare?ref=badge_large)