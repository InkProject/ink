title: "An Elegant Static Blog Generator —— InkPaper"
date: 2015-03-01 17:00:00 +0800
update: 2016-07-11 17:00:00 +0800
author: me
cover: "-/images/example-en.png"
tags:
    - 设计
    - 写作
preview: InkPaper is an static blog generator developed by Golang, No dependencies, Cross platform, Easy use, Fast build, Elegant theme.

---

## Introduce

InkPaper is an static blog generator developed by Golang, No dependencies, Cross platform, Easy use, Fast build, Elegant theme.

![InkPaper - An Elegant Static Blog Generator](-/images/example-en.png)

### Quick Start
- Download & Extract [Ink](http://www.chole.io/)，Run `ink preview`

  > Tips：Linux/macOS, use `./ink preview`

- Open `http://localhost:8000` in browser to preview

### Website Configuration
Edit `config.yml`, use format:

``` yaml
site:
    title: Website Title
    subtitle: Website Subtitle
    limit: Max Article Count Per Page
    theme: Website Theme Directory
    comment: Comment Plugin Variable (Default is disqus username)
    root: Website Root Path #Optional
    lang: Website Language #Support en, zh, ru, ja, Configurable in theme/lang.yml
    url: Website URL #For RSS Generating
    link: Article Link Scheme #Default Is {title}.html，Support {year},{month},{day},{title} Variables

authors:
    AuthorID:
        name: Author Name
        intro: Author Motto
        avatar: Author Avatar Path

build:
    output: Build Output Directory #Optional, Default is "public"
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
draft: false #If Draft，Optional
top: false #Place article to top, Optional
preview: Article Preview，Also use <!--more--> to split in body #Optional
tags: #Optional
    - Tag1
    - Tag2
type: post #Specify type is post or page, Optional
hide: false #Hide article，can be accessed via URL, Optional

---

Markdown Format's Body
```

### Publish
- Run `ink publish` in blog directory to automatically build and publish
- Or run `ink build` to manually deploy generated `public` directory

> **Tips**: When files changed，`ink preview` will automatically rebuild blog，refresh browser to update.

## Customization

### Modify Theme

Default theme placed in `theme` folder, run `npm install` and `npm run build` to rebuild in this folder.

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

## Theme

- Dark(Official Theme): [https://github.com/InkProject/ink-theme-dark](https://github.com/InkProject/ink-theme-dark)
- simple: [https://github.com/myiq/ink-simple](https://github.com/myiq/ink-simple)

## Related Toturials

- [Automatically deploy your Ink blog to GitHub pages wiht Travis CI](http://www.shery.me/blog/travis-ci.html)

## Issue Report

[CC Attribution-NonCommercial License 4.0](https://creativecommons.org/licenses/by-nc/4.0/)

[https://github.com/InkProject/ink/issues](https://github.com/InkProject/ink/issues)

## Change Log

- [2016-07-11] Fix bugs and merge many PRs
- [2015-08-15] Bug fix, support RSS feed, improve theme
- [2015-07-04] Bug fix, improve theme, support top, i18n, subtemplate
- [2015-06-04] Build more platform, add archive and tag page
- [2015-03-01] Release first beta version

## Develop Plan

- Improve Theme
- InkPaper Editor

## They are using

- [http://www.chole.io/blog/](http://www.chole.io/blog/)
- [http://blog.hyper.sh/](http://blog.hyper.sh/)
- [http://wangxu.me/](http://wangxu.me/)
- [http://whzecomjm.com/](http://whzecomjm.com/)
- [http://www.shery.me/blog/](http://www.shery.me/blog/)
