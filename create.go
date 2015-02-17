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
    var tagMap = make(map[string]Articles)
    // Parse config
    globalConfig := ParseGlobalConfig(filepath.Join(root, "config.yml"))
    themePath := filepath.Join(root, globalConfig.Site.Theme)
    publicPath := filepath.Join(root, "public")
    sourcePath := filepath.Join(root, "source")
    // Compile template
    var articleTpl = CompileTpl(filepath.Join(themePath, "article.html"), "article")
    var pageTpl = CompileTpl(filepath.Join(themePath, "page.html"), "page")
    // Clean public folder
    os.RemoveAll(publicPath)
    // Find all .md to generate article
    filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
        fileExt := strings.ToLower(filepath.Ext(path))
        if fileExt == ".md" {
            // Parse markdown data
            article := ParseMarkdown(path)
            // Generate page name
            fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
            // Generate directory
            directory := time.Unix(article.Date, 0).Format("2006/01/02/")
            err := os.MkdirAll(filepath.Join(publicPath, directory), 0777)
            if err != nil {
                Fatal(err.Error())
            }
            outPath := directory + fileName + ".html"
            // Generate file path
            article.Link = outPath
            article.GlobalConfig = *globalConfig
            articles = append(articles, *article)
            // Get tags info
            for _, tag := range article.Tag {
                if _, ok := tagMap[tag]; !ok {
                    tagMap[tag] = make(Articles, 0)
                }
                tagMap[tag] = append(tagMap[tag], *article)
            }
            // Render article
            RenderPage(articleTpl, article, filepath.Join(publicPath, outPath))
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
        outPath := filepath.Join(publicPath, "index.html")
        if i != 0 {
            fileName := "page" + strconv.Itoa(i + 1) + ".html"
            outPath = filepath.Join(publicPath, fileName)
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
            "TagName": "",
        }
        RenderPage(pageTpl, data, outPath)
    }
    // Copy static files
    CopyDir(filepath.Join(themePath, "css"), filepath.Join(publicPath, "css"))
    CopyDir(filepath.Join(themePath, "js"), filepath.Join(publicPath, "js"))
    CopyDir(filepath.Join(sourcePath, "image"), filepath.Join(publicPath, "image"))
}
