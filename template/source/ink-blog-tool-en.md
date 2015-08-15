title: An Elegant Static Blog Generator —— InkPaper
date: 2015-03-01 19:00:00 +0800
update: 2015-07-04 10:00:00 +0800
author: me
cover: -/images/example-en.png
tags:
    - Design
    - Writing
preview: InkPaper is an static blog generator developed by Golang, No dependencies, Cross platform, Easy use, Fast build, Elegant theme

---

## Introduce

InkPaper is an static blog generator developed by Golang, No dependencies, Cross platform, Easy use, Fast build, Elegant theme.

![InkPaper - An Elegant Static Blog Generator](-/images/example-en.png)

### Quick Start
- Download & Extract [Ink](http://www.inkpaper.io/)，Run `ink preview`
- Open `http://localhost:8000` in browser to preview

### Website Configuration
Edit `config.yml`, use format:

``` yaml
site:
    title: Website Title
    subtitle: Website Subtitle
    limit: Max Article Count Per Page
    theme: Website Theme Directory
    disqus: Disqus Plugin Username
    root: Website Root Path #Optional
    lang: Website language #Support en, zh, Configurable in theme/lang.yml

authors:
    AuthorID:
        name: Author Name
        intro: Author Motto
        avatar: Author Avatar Path

build:
    port: Preview Port
    copy:
        - Copied Files When Build
    publish: |
        Excuted command when use 'ink publish'
```

### Writing
Create any `.md` file in `source` directory (Support subdirectory), use format:

``` yaml
title: Article Title
date: Year-Month-Day Hour:Minute:Second #Created Time，Support TimeZone, such as " +0800"
update: Year-Month-Day Hour:Minute:Second #Updated Time，Optional，Support TimeZone, such as " +0800"
author: AuthorID
cover: Article Cover Path #Optional
draft: true #If Draft，Optional
top: Place article to top #Optional
preview: Article Preview，Also use <!--more--> to split in body #Optional
tag: #Optional
    - Tag1
    - Tag2

---

Markdown Format's Body
```

### Publish
- Run `ink publish` in blog directory to automatically build and publish
- Or run `ink build` to manually deploy generated `public` directory

> **Tips**: When `source` directory changed，`ink preview` will automatically rebuild blog，refresh browser to update

## Customization

### Modify Theme

Default theme use coffee & less build, after modify that files, run `gulp` in `theme` to recompile, run `ink` will copy js and css directory to public directory;

page `page.html` (article list) and `article.html` (article), use variable with [Golang Template](http://golang.org/pkg/html/template/) syntax.

### New Page

Created any `.html` file will be copied to `source` directory, could use all variables on `site` field in `config.yml`.

### Blog Migrate (Beta)

Support simple Jeklly/Hexo post convert, use:

``` shell
ink convert /path/_posts
```

### Build from source

Local Build

1. Install [Golang](http://golang.org/doc/install) environment
2. Run `go get github.com/InkProject/ink`, compile and get ink
3. Run `ink preview $GOPATH/src/github.com/InkProject/ink/template`, preview blog

Docker Build (Example)

1. Clone code `git clone git@github.com:InkProject/ink.git`
2. Build image `docker build -t ink .` in source directory
3. Run container `docker run -p 8888:80 ink`

## License
[CC Attribution-NonCommercial License 4.0](https://creativecommons.org/licenses/by-nc/4.0/)

## Issue Report

[https://github.com/InkProject/ink/issues](https://github.com/InkProject/ink/issues)

## Change Log

- [2015-08-15] Bug fix, support RSS feed, improve theme
- [2015-07-04] Bug fix, improve theme, support top, i18n, subtemplate
- [2015-06-04] Build more platform, add archive and tag page
- [2015-03-01] Release first beta version

## Develop Plan

- Improve Theme
- Support RSS Feed
- Extension And Plugin

## They are using

- [http://www.inkpaper.io/blog/](http://www.inkpaper.io/blog/)
- [https://hyper.sh/blog/](https://hyper.sh/blog/)
- [http://wangsiyi.net/](http://wangsiyi.net/)
- [http://lubia.cn/](http://lubia.cn/)
