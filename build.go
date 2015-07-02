package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"
)

// Parse config
var articleTpl, pageTpl, archiveTpl, tagTpl template.Template
var themePath, publicPath, sourcePath string

// For concurrency
var wg sync.WaitGroup

// Data struct
type ArticleInfo struct {
	Date  string
	Title string
	Link  string
}

type Archive struct {
	Year     string
	Articles Collections
}

type Tag struct {
	Name     string
	Count    int
	Articles Collections
}

// For sort
type Collections []interface{}

func (v Collections) Len() int      { return len(v) }
func (v Collections) Swap(i, j int) { v[i], v[j] = v[j], v[i] }
func (v Collections) Less(i, j int) bool {
	switch v[i].(type) {
	case ArticleInfo:
		return v[i].(ArticleInfo).Date > v[j].(ArticleInfo).Date
	case Article:
		return v[i].(Article).Date > v[j].(Article).Date
	case Archive:
		return v[i].(Archive).Year > v[j].(Archive).Year
	case Tag:
		return v[i].(Tag).Count > v[j].(Tag).Count
	}
	return false
}

func Build() {
	startTime := time.Now()
	var articles = make(Collections, 0)
	var tagMap = make(map[string]Collections)
	var archiveMap = make(map[string]Collections)
	// Parse config
	themePath = filepath.Join(rootPath, globalConfig.Site.Theme)
	publicPath = filepath.Join(rootPath, "public")
	sourcePath = filepath.Join(rootPath, "source")
	// Compile template
	articleTpl = CompileTpl(filepath.Join(themePath, "article.html"), "article")
	pageTpl = CompileTpl(filepath.Join(themePath, "page.html"), "page")
	archiveTpl = CompileTpl(filepath.Join(themePath, "archive.html"), "archive")
	tagTpl = CompileTpl(filepath.Join(themePath, "tag.html"), "tag")
	// Clean public folder
	cleanPatterns := []string{"post", "tag", "images", "js", "css", "*.html", "favicon.ico", "robots.txt"}
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
			article, err := ParseMarkdown(path)
			if err != nil {
				Log(err)
				return err
			}
			if article.Draft {
				return nil
			}
			// Generate page name
			fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
			Log("Building " + fileName)
			// Generate directory
			unixTime := time.Unix(article.Date, 0)
			directory := unixTime.Format("post/2006/01/02/")
			err = os.MkdirAll(filepath.Join(publicPath, directory), 0777)
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
					tagMap[tag] = make(Collections, 0)
				}
				tagMap[tag] = append(tagMap[tag], *article)
			}
			// Get archive info
			dateYear := unixTime.Format("2006")
			if _, ok := archiveMap[dateYear]; !ok {
				archiveMap[dateYear] = make(Collections, 0)
			}
			articleInfo := ArticleInfo{
				Date:  unixTime.Format("2006-01-02"),
				Title: article.Title,
				Link:  article.Link,
			}
			archiveMap[dateYear] = append(archiveMap[dateYear], articleInfo)
			// Render article
			wg.Add(1)
			go RenderPage(articleTpl, article, filepath.Join(publicPath, outPath))
		}
		return nil
	})
	// Sort by date
	sort.Sort(articles)
	// Generate article pages
	wg.Add(1)
	go RenderArticles("", articles, "")
	// Generate tags pages
	for tagName, articles := range tagMap {
		wg.Add(1)
		go RenderArticles(filepath.Join("tag", tagName), articles, tagName)
	}
	// Generate archive page
	archives := make(Collections, 0)
	for year, articleInfos := range archiveMap {
		// Sort by date
		sort.Sort(articleInfos)
		archives = append(archives, Archive{
			Year:     year,
			Articles: articleInfos,
		})
	}
	// Sort by year
	sort.Sort(archives)
	wg.Add(1)
	go RenderPage(archiveTpl, map[string]interface{}{
		"Total":   len(articles),
		"Archive": archives,
		"Site":    globalConfig.Site,
	}, filepath.Join(publicPath, "archive.html"))
	// Generate tag page
	tags := make(Collections, 0)
	for tagName, tagArticles := range tagMap {
		articleInfos := make(Collections, 0)
		for _, article := range tagArticles {
			articleInfos = append(articleInfos, ArticleInfo{
				Date:  time.Unix(article.(Article).Date, 0).Format("2006-01-02"),
				Title: article.(Article).Title,
				Link:  article.(Article).Link,
			})
		}
		// Sort by date
		sort.Sort(articleInfos)
		tags = append(tags, Tag{
			Name:     tagName,
			Count:    len(tagArticles),
			Articles: articleInfos,
		})
		// Sort by count
		sort.Sort(Collections(tags))
	}
	wg.Add(1)
	go RenderPage(tagTpl, map[string]interface{}{
		"Total": len(articles),
		"Tag":   tags,
		"Site":  globalConfig.Site,
	}, filepath.Join(publicPath, "tag.html"))
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
	Log("Copying files")
	Copy()
	wg.Wait()
	endTime := time.Now()
	usedTime := endTime.Sub(startTime)
	fmt.Printf("\nBuild finish in public folder (%v)\n", usedTime)
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
