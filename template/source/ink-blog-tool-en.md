title: "An Elegant Static Blog Generator —— InkPaper"
date: 2015-03-01 17:00:00 +0800
update: 2015-11-06 10:00:00 +0800
author: me
cover: "-/images/example-en.png"
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
    comment: Comment Plugin Variable (Default Is Disqus Username)
    root: Website Root Path #Optional
    lang: Website Language #Support en, zh, Configurable in theme/lang.yml
    url: Website URL #For RSS Generating
    link: Article Link Scheme #Default Is {title}.html，Support {year},{month},{day},{title} Variables

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
draft: false #If Draft，Optional
top: false #Place article to top, Optional
preview: Article Preview，Also use <!--more--> to split in body #Optional
tags: #Optional
    - Tag1
    - Tag2
