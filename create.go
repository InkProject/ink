package main

import (
    "os"
    "sort"
    "time"
    "strconv"
    "strings"
    "path/filepath"
)

func Create(root string) {
    var articles = make(Articles, 0)
    // Compile template
    var articleTpl = CompileTpl(root + "/article.html", "article")
    var pageTpl = CompileTpl(root + "/page.html", "page")
    // Parse config
    globalConfig := ParseGlobalConfig(root + "/config.yml")
    // Clean public folder
    os.RemoveAll(root + "/public")
    // Find all .md to generate article
    filepath.Walk(root + "/source", func(path string, info os.FileInfo, err error) error {
        fileExt := strings.ToLower(filepath.Ext(path))
        if fileExt == ".md" {
            // Parse markdown data
            article := ParseMarkdown(path)
            // Generate page name
            fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
            // Generate directory
            directory := time.Unix(article.Date, 0).Format("2006/01/02/")
            err := os.MkdirAll(root + "/public/" + directory, 0777)
            if err != nil {
                Fatal(err.Error())
            }
            outPath := directory + fileName + ".html"
            // Generate file path
            article.Link = outPath
            article.GlobalConfig = *globalConfig
            articles = append(articles, *article)
            // Render article
            RenderPage(articleTpl, article, root + "/public/" + outPath)
        }
        return nil
    })
    // Sort by time
    sort.Sort(Articles(articles))
    // Split page
    limit := globalConfig.Site.Limit
    total := len(articles)
    page := total / limit
    rest := total % limit
    if rest != 0 {
        page++
    }
    if total < limit {
        page = 1
    }
    for i := 0; i < page; i ++ {
        var prev = "/page" + strconv.Itoa(i) + ".html"
        var next = "/page" + strconv.Itoa(i + 2) + ".html"
        outPath := root + "/public/index.html"
        if i != 0 {
            outPath = root + "/public/page" + strconv.Itoa(i + 1) + ".html"
        } else {
            prev = ""
        }
        if i == 1 {
            prev = "/index.html"
        }
        first := i * limit
        count := first + limit
        if i == page - 1 {
            if rest != 0 {
                count = first + rest
            }
            next = ""
        }
        var data = map[string]interface{}{
            "Articles": articles[first:count],
            "Site": globalConfig.Site,
            "Author": globalConfig.Author,
            "Page": i + 1,
            "Total": page,
            "Prev": prev,
            "Next": next,
        }
        RenderPage(pageTpl, data, outPath)
    }
    // Copy static files
    CopyDir(root + "/css", root + "/public/css")
    CopyDir(root + "/js", root + "/public/js")
    CopyDir(root + "/source/image", root + "/public/image")
}
