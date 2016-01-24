title: "简洁的静态博客构建工具 —— 纸小墨（InkPaper）"
date: 2015-03-01 18:00:00 +0800
update: 2015-11-06 10:00:00 +0800
author: me
cover: "-/images/example.png"
tags:
    - 设计
    - 写作
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
    comment: 评论插件变量(默认为Disqus账户名)
    root: 网站根路径 #可选
    lang: 网站语言 #支持en, zh，可在theme/lang.yml配置
    url: 网站链接 #用于RSS生成
    link: 文章链接形式 #默认为{title}.html，支持{year},{month},{day},{title}变量

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
draft: false #草稿，可选
top: false #置顶文章，可选
preview: 文章预览，也可在正文中使用<!--more-->分割 #可选
tags: #可选
    - 标签1
    - 标签2