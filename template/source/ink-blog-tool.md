title: 简洁的静态博客构建工具 —— 纸小墨（InkPaper）
date: 2015-03-01 18:00:00 +0800
update: 2015-07-04 10:00:00 +0800
author: me
cover: -/images/example.png
tags:
    - Design
    - Writing
preview: 纸小墨（InkPaper）是一个使用GO语言编写的静态博客构建工具，可以快速搭建博客网站。优点是无依赖跨平台，配置简单构建快速，注重简洁易用与排版优化

---

## 纸小墨简介

纸小墨（InkPaper）是一个使用GO语言编写的静态博客构建工具，可以快速搭建博客网站。优点是无依赖跨平台，配置简单构建快速，注重于简洁易用与排版优化。

![纸小墨 - 简洁的静态博客构建工具](-/images/example.png)

### 开始上手
- 下载并解压 [Ink](http://www.inkpaper.io/)，运行命令 `ink preview`
- 使用浏览器访问 `http://localhost:8000` 预览

### 配置网站
编辑`config.yml`，使用如下格式

``` yaml
site:
    title: 网站标题
    subtitle: 网站子标题
    limit: 每页可显示的文章数目
    theme: 网站主题目录
    disqus: Disqus评论插件账户名
    root: 网站根路径 #可选
    lang: 网站语言 #支持en, zh，可在theme/lang.yml配置
authors:
    作者ID:
        name: 作者名称
        intro: 作者简介
        avatar: 作者头像路径
build:
    port: 预览端口
    copy:
        - 构建时将会复制的目录/文件
    publish: |
        ink publish 命令将会执行的脚本
```

### 创建文章
在`source`目录中建立任意`.md`文件（可置于子文件夹），使用如下格式

``` yaml
title: 文章标题
date: 年-月-日 时:分:秒 #创建时间，可加时区如" +0800"
update: 年-月-日 时:分:秒 #更新时间，可选，可加时区如" +0800"
author: 作者ID
cover: 题图链接 #可选
draft: true #草稿，可选
top: true #置顶文章，可选
preview: 文章预览，也可在正文中使用<!--more-->分割 #可选
tag: #可选
    - 标签1
    - 标签2

---

Markdown格式的正文
```

### 发布博客
- 在博客目录下运行`ink publish`命令自动构建博客并发布
- 或运行`ink build`命令将生成的`public`目录下的内容手动部署

> Tips: 当使用`ink preview`命令时，`source`目录中文件发生变化，会自动重新构建博客，刷新浏览器以更新

## 定制支持

### 修改主题

默认主题使用coffee, less编写，修改对应文件后，需要在`theme`目录下运行`gulp`命令重新编译，使用`ink build`命令构建时默认将会复制js, css目录到`public`目录下。

页面包含`page.html`（文章列表）及`article.html`（文章）等，所有页面均支持[GO语言HTML模板](http://golang.org/pkg/html/template/)语法，可引用变量。

### 添加页面

在`source`目录下创建的任意`.html`文件将被复制，这些文件中可引用`config.yml`中site字段下的所有变量。

### 博客迁移(Beta)

Ink提供简单的Jeklly/Hexo博客文章格式转换，使用命令：
``` shell
ink convert /path/_posts
```

### 源码编译

本地运行

1. 配置[GO](http://golang.org/doc/install)语言环境
2. 运行命令`go get github.com/InkProject/ink`，编译并获取ink
3. 运行命令`ink preview $GOPATH/src/github.com/InkProject/ink/template`，预览博客

Docker构建（示例）

1. Clone源码 `git clone git@github.com:InkProject/ink.git`
2. 源码目录下构建镜像`docker build -t ink .`
3. 运行容器`docker run -p 8000:80 ink`

## 开源协议
[CC Attribution-NonCommercial License 4.0](https://creativecommons.org/licenses/by-nc/4.0/)

## 反馈建议

请报告至 [https://github.com/InkProject/ink/issues](https://github.com/InkProject/ink/issues)

## 更新历史

- [2015-08-15] 一些Bug修复，新增RSS订阅支持，主题改善
- [2015-07-04] 数项Bug修复与主题改善，支持置顶，多语言，子模板
- [2015-06-04] 编译更多平台版本，增加标签与存档页
- [2015-03-01] Beta版本，基础功能完成

## 更新计划

- 排版深度优化
- 页面SEO优化
- 扩展与插件支持

## 他们正在使用

- [http://www.inkpaper.io/blog/](http://www.inkpaper.io/blog/)
- [https://hyper.sh/blog/](https://hyper.sh/blog/)
- [http://wangsiyi.net/](http://wangsiyi.net/)
- [http://lubia.cn/](http://lubia.cn/)
