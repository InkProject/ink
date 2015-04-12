package main

import (
    "os"
    "fmt"
    "time"
    "sort"
    "sync"
    "strconv"
    "strings"
    "html/template"
    "path/filepath"
)

// Parse config
var pageTpl template.Template
var themePath, publicPath, sourcePath string

// For concurrency
var wg sync.WaitGroup

func Build() {
    startTime := time.Now()
    var articles = make(Articles, 0)
    var tagMap = make(map[string]Articles)
    // Parse config
    themePath = filepath.Join(rootPath, globalConfig.Site.Theme)
    publicPath = filepath.Join(rootPath, "public")
    sourcePath = filepath.Join(rootPath, "source")
    // Compile template
    articleTpl := CompileTpl(filepath.Join(themePath, "article.html"), "article")
    pageTpl = CompileTpl(filepath.Join(themePath, "page.html"), "page")
    // Clean public folder
    cleanPatterns := []string{"post", "tag", "images", "js", "css", "*.html"}
    for _, pattern := range cleanPatterns {
        files, _ := filepath.Glob(filepath.Join(publicPath, pattern))
        for _, path := range files {
            os.RemoveAll(path)
        }
    }
    // Find all .md to generate article
    filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
        fileExt := strings.ToLower(filepath.Ext(path))
        if fileExt == ".md" {
            // Parse markdown data
            article := ParseMarkdown(path)
            if article.Draft {
                return nil
            }
            // Generate page name
            fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
            Log("Building " + fileName)
            // Generate directory
            directory := time.Unix(article.Date, 0).Format("post/2006/01/02/")
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
            for _, tag := range article.Tags {
                if _, ok := tagMap[tag]; !ok {
                    tagMap[tag] = make(Articles, 0)
                }
                tagMap[tag] = append(tagMap[tag], *article)
            }
            // Render article
            wg.Add(1)
            go RenderPage(articleTpl, article, filepath.Join(publicPath, outPath))
        }
        return nil
    })
    // Generate article pages
    wg.Add(1)
    go RenderArticles("", articles, "")
    // Generate tags pages
    for tagName, articles := range tagMap {
        wg.Add(1)
        go RenderArticles(filepath.Join("tag", tagName), articles, tagName)
    }
    // Generate other pages
    files, _ := filepath.Glob(filepath.Join(sourcePath, "*.html"))
    for _, path := range files {
        fileExt := strings.ToLower(filepath.Ext(path))
        baseName := filepath.Base(path)
        if fileExt == ".html" {
            htmlTpl := CompileTpl(path, baseName)
            relPath, _ := filepath.Rel(sourcePath, path)
            wg.Add(1)
            go RenderPage(htmlTpl, globalConfig, filepath.Join(publicPath, relPath))
        }
    }
    // Copy static files
    Copy()
    wg.Wait()
    endTime := time.Now()
    usedTime := endTime.Sub(startTime)
    fmt.Printf("\nBuild finish in public folder (%v)\n", usedTime)
}

// Generate html file by article data
func RenderArticles(rootPath string, articles Articles, tagName string) {
    defer wg.Done()
    // Create path
    pagePath := filepath.Join(publicPath, rootPath)
    os.MkdirAll(pagePath, 0777)
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
        var prev = filepath.Join(rootPath, "page" + strconv.Itoa(i) + ".html")
        var next = filepath.Join(rootPath, "page" + strconv.Itoa(i + 2) + ".html")
        outPath := filepath.Join(pagePath, "index.html")
        if i != 0 {
            fileName := "page" + strconv.Itoa(i + 1) + ".html"
            outPath = filepath.Join(pagePath, fileName)
        } else {
            prev = ""
        }
        if i == 1 {
            prev = filepath.Join(rootPath, "index.html")
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
            "Page": i + 1,
            "Total": page,
            "Prev": prev,
            "Next": next,
            "TagName": tagName,
            "TagCount": len(articles),
        }
        wg.Add(1)
        go RenderPage(pageTpl, data, outPath)
    }
}

// Copy static files
func Copy() {
    srcList := globalConfig.Build.Copy
    for _, source := range srcList {
        srcPath := filepath.Join(rootPath, source)
        file, err := os.Stat(srcPath)
        if err != nil {
            Fatal("Not exist: " + srcPath)
        }
        fileName := file.Name()
        desPath := filepath.Join(publicPath, fileName)
        wg.Add(1)
        if file.IsDir() {
            go CopyDir(srcPath, desPath)
        } else {
            go CopyFile(srcPath, desPath)
        }
    }
}
