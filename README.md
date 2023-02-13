## Introduction

InkPaper is a static blog generator developed in Golang. No dependencies, cross platform, easy to use, fast building times and an elegant theme.

[![apm](https://img.shields.io/badge/license-CC%20BY--NC%204.0-blue.svg)](https://creativecommons.org/licenses/by-nc/4.0/)

![InkPaper - An Elegant Static Blog Generator](template/source/images/example-en.png)

### Features
- YAML format configuration
- Markdown format articles
- No dependencies, cross platform
- Super fast build times
- Continuously improving theme and typography
- Multiple article authors support
- Archive and tag generation
- Real-time preview when saving
- Offline full-text keyword search
- $\LaTeX$ style math formula support (MathJax):

$$
\int_{-\infty}^\infty g(x) dx = \frac{1}{2\pi i} \oint_{\gamma} \frac{f(z)}{z-g(x)} dz
$$

### Quick Start
- Download & Extract [Ink](https://github.com/InkProject/ink/releases) and run `ink preview`

  > Tip：Linux/macOS, use `./ink preview`

- Open `http://localhost:8000` in your browser to preview

### Website Configuration
Edit `config.yml`, use this format:

``` yaml
site:
    title: Website Title
    subtitle: Website Subtitle
    limit: Max Article Count Per Page
    theme: Website Theme Directory
    comment: Comment Plugin Variable (Default is disqus username)
    root: Website Root Path # Optional
    lang: Website Language # Support en, zh, ru, ja, de, pt-br, configurable in theme/lang.yml
    url: Website URL # For RSS generating
    link: Article Link Scheme # Default is {title}.html, Support {year}, {month}, {day}, {hour}, {minute}, {second}, {title} variables

authors:
    AuthorID: # Your author ID, used in article's author field
        name: Author Name
        intro: Author Motto
        avatar: Author Avatar Path

build:
    output: Build Output Directory # Optional, default is "public"
    port: Preview Port
    copy:
        - Copied Files When Build
    publish: |
        Excuted command when 'ink publish' is used
```

### Blog Writing
Create a `.md` file in the `source` directory (Supports subdirectories). Use this format:

``` yaml
title: Article Title
date: Year-Month-Day Hour:Minute:Second #Created Time. Support timezone, such as " +0800"
update: Year-Month-Day Hour:Minute:Second #Updated Time, optional. Support timezone, such as " +0800"
author: AuthorID
cover: Article Cover Path # Optional
draft: false # Is draft or not, optional
top: false # Place article to top or not, optional
preview: Article Preview, Also use <!--more--> to split in body # Optional
tags: # Optional
    - Tag1
    - Tag2
type: post # Specify type is post or page, optional
hide: false # Hide article or not. Hidden atricles still can be accessed via URL, optional
toc: false # Show table of contents or not, optional
---

Markdown Format's Body
```

### Publish
- Run `ink publish` in the blog directory to automatically build and publish
- Or run `ink build` to manually deploy generated `public` directory

> **Tips**: When files changed, `ink preview` will automatically rebuild the blog. Refresh browser to update.

## Customization

### Modifying The Theme

The default theme is placed in the `theme` folder, run `npm install` and `npm run build` to rebuild in this folder.

page `page.html` (article list) and `article.html` (article), use variable with [Golang Template](http://golang.org/pkg/html/template/) syntax.

### New Page

Created any `.html` file will be copied to `source` directory, could use all variables on `site` field in `config.yml`.


#### Define Custom Variables
InkPaper supports defining custom variables in pages, which must be placed under `site.config` in `config.yaml`, such as:

``` yaml
site:
    config:
        MyVar: "Hello World"
```

The variable can be referenced in the page by `{{.Site.Config.MyVar}}`.

> **Note**
>
> Although the field names in other parts of `config.yaml` are all lowercase, the name of the custom variable must be used correctly. Otherwises, such a variable:
>
> ```yaml
> site:
>     config:
>         MYVAR_aAa: "Hello World"
> ```
> must be referenced in the page as `{{.Site.Config.MYVAR_aAa}}`.

#### Use Functions (Experimental)

InkPaper defines a minimal set of functions that can be used in HTML pages (except for `.md` source files), such as

``` yaml
{{ readFile "path/to/file" }}
```

This will read the content of the file `path/to/file` and include it in the page without any processing.

For file-related functions, when executed in the `source` directory, the file path is relative to the `source` directory; when executed in other directories, the file path is relative to the theme (such as `theme`).

See the source file `funcs.go` for a list of all functions.

### Blog Migration (Beta)

Supports simple Jeklly/Hexo post convertions. Usage:

``` shell
ink convert /path/_posts
```

### Building from source

**Local Build**

1. Install [Golang](http://golang.org/doc/install) environment
2. Run `git clone https://github.com/InkProject/ink && cd ink && go install` to compile and install ink
3. Run `ink preview $GOPATH/src/github.com/InkProject/ink/template` to preview blog

**Docker Build (Example)**

1. Clone code `git clone git@github.com:InkProject/ink.git`
2. Build image `docker build -t ink .` in source directory
3. Run container `docker run -p 8888:80 ink`

## Theme

- Dark (Official Theme): [https://github.com/InkProject/ink-theme-dark](https://github.com/InkProject/ink-theme-dark)
- Simple: [https://github.com/myiq/ink-simple](https://github.com/myiq/ink-simple)
- Story: [https://github.com/akkuman/ink-theme-story](https://github.com/akkuman/ink-theme-story)
- Material2: [https://github.com/w568w/InkMaterialTheme](https://github.com/w568w/InkMaterialTheme)

## Related Tutorials

- [Automatically deploy your Ink blog to GitHub pages with Travis CI](http://www.shery.me/blog/travis-ci.html)

## License
[CC Attribution-NonCommercial License 4.0](https://creativecommons.org/licenses/by-nc/4.0/)

## Reporting An Issue

[https://github.com/InkProject/ink/issues](https://github.com/InkProject/ink/issues)

## These blogs are driven by InkPaper

- [https://imeoer.github.io/blog/](https://imeoer.github.io/blog/)
- [http://blog.hyper.sh/](http://blog.hyper.sh/)
- [http://wangxu.me/](http://wangxu.me/)
- [http://whzecomjm.com/](http://whzecomjm.com/)
- [http://www.shery.me/blog/](http://www.shery.me/blog/)
