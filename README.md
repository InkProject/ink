## Ink 简介

Ink 是一个使用GO语言编写的静态博客生成器，使用它可以快速构建博客网站。无依赖跨平台，配置简单，构建快速，支持多用户，默认主题简洁，优化了中文排版，提高阅读体验。

### 开始上手
- 下载 [Ink 工具](https://github.com/InkProject/ink/releases)，然后下载并解压 [快速开始模板](https://github.com/InkProject/blog)
- 在模板目录下运行命令 `ink preview`
- 使用浏览器访问 `http://localhost:8888` 预览

### 配置网站
编辑`config.yml`，使用如下格式

``` yaml
site:
    title: 网站标题
    subtitle: 网站子标题
    limit: 每页可显示的文章数目
    theme: 网站主题目录
    disqus: Disqus评论插件账户名

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
在`source`目录中建立任意`.md`文件，使用如下格式

``` yaml
title: 文章标题
date: 年-月-日 时:分:秒
author: 作者ID
topic: 题图链接 #可选
preview: 文章预览，也可在正文中使用--more--分割 #可选
tag: 空格 分割 标签 #可选

---

Markdown格式的正文
```

### 发布博客
- 在博客目录下运行命令`ink publish`构建博客并发布
- 也可以运行`ink`命令将生成的`public`目录下的内容部署到服务器

## Markdown 样式支持

### 图片

![图片说明](/image/example.jpg)

### 引用

> Markdown 是一种轻量级标记语言，创始人为約翰·格魯伯（John Gruber）。它允许人们“使用易读易写的纯文本格式编写文档，然后转换成有效的XHTML(或者HTML)文档”。这种语言吸收了很多在电子邮件中已有的纯文本标记的特性。
—— [维基百科](http://www.wikiwand.com/zh/Markdown)

### 代码
``` python
@requires_authorization
def somefunc(param1='', param2=0):
    r'''A docstring'''
    if param1 > param2: # interesting
        print 'Gre\'ater'
    return (param2 - param1 + 1) or None

class SomeClass:
    pass

>>> message = '''interpreter
... prompt'''
```

### 表格
| 左对齐    |    右对齐| 居中 |
| :-------- | -------:| :--: |
| apple     |     100 |  1   |
| banana    |     200 |  2   |
| pear      |     300 |  3   |
