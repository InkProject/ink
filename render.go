package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/feeds"
)

type Data interface{}

type RenderArticle struct {
	Article
	Next *Article
	Prev *Article
}

// Compile html template
func CompileTpl(tplPath string, partialTpl string, name string) template.Template {
	// Read template data from file
	html, err := ioutil.ReadFile(tplPath)
	if err != nil {
		Fatal(err.Error())
	}
	// Append partial template
	htmlStr := string(html) + partialTpl
	funcMap := template.FuncMap{
		"i18n": func(val string) string {
			return globalConfig.I18n[val]
		},
	}
	// Generate html content
	tpl, err := template.New(name).Funcs(funcMap).Parse(htmlStr)
	if err != nil {
		Fatal(err.Error())
	}
	return *tpl
}

// Render html file by data
func RenderPage(tpl template.Template, tplData interface{}, outPath string) {
	// Create file
	outFile, err := os.Create(outPath)
	if err != nil {
		Fatal(err.Error())
	}
	defer func() {
		outFile.Close()
	}()
	defer wg.Done()
	// Template render
	err = tpl.Execute(outFile, tplData)
	if err != nil {
		Fatal(err.Error())
	}
}

// Generate all article page
func RenderArticles(tpl template.Template, articles Collections) {
	defer wg.Done()
	articleCount := len(articles)
	for i, _ := range articles {
		currentArticle := articles[i].(Article)
		var renderArticle = RenderArticle{currentArticle, nil, nil}
		if i >= 1 {
			article := articles[i-1].(Article)
			renderArticle.Prev = &article
		}
		if i <= articleCount-2 {
			article := articles[i+1].(Article)
			renderArticle.Next = &article
		}
		outPath := filepath.Join(publicPath, currentArticle.Link)
		wg.Add(1)
		go RenderPage(tpl, renderArticle, outPath)
	}
}

// Generate rss page
func GenerateRSS(articles Collections) {
	defer wg.Done()
	var feedArticles Collections
	if len(articles) < globalConfig.Site.Limit {
		feedArticles = articles
	} else {
		feedArticles = articles[0:globalConfig.Site.Limit]
	}
	if globalConfig.Site.Url != "" {
		feed := &feeds.Feed{
			Title:       globalConfig.Site.Title,
			Link:        &feeds.Link{Href: globalConfig.Site.Url},
			Description: globalConfig.Site.Subtitle,
			Author:      &feeds.Author{globalConfig.Site.Title, ""},
			Created:     time.Now(),
		}
		feed.Items = make([]*feeds.Item, 0)
		for _, item := range feedArticles {
			article := item.(Article)
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       article.Title,
				Link:        &feeds.Link{Href: globalConfig.Site.Url + "/" + article.Link},
				Description: string(article.Preview),
				Author:      &feeds.Author{article.Author.Name, ""},
				Created:     article.Time,
				Updated:     article.MTime,
			})
		}
		if atom, err := feed.ToAtom(); err == nil {
			err := ioutil.WriteFile(filepath.Join(publicPath, "atom.xml"), []byte(atom), 0644)
			if err != nil {
				Fatal(err.Error())
			}
		} else {
			Fatal(err.Error())
		}
	}
}

// Generate article list page
func RenderArticleList(rootPath string, articles Collections, tagName string) {
	defer wg.Done()
	// Create path
	pagePath := filepath.Join(publicPath, rootPath)
	os.MkdirAll(pagePath, 0777)
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
	for i := 0; i < page; i++ {
		var prev = filepath.Join(rootPath, "page"+strconv.Itoa(i)+".html")
		var next = filepath.Join(rootPath, "page"+strconv.Itoa(i+2)+".html")
		outPath := filepath.Join(pagePath, "index.html")
		if i != 0 {
			fileName := "page" + strconv.Itoa(i+1) + ".html"
			outPath = filepath.Join(pagePath, fileName)
		} else {
			prev = ""
		}
		if i == 1 {
			prev = filepath.Join(rootPath, "index.html")
		}
		first := i * limit
		count := first + limit
		if i == page-1 {
			if rest != 0 {
				count = first + rest
			}
			next = ""
		}
		var data = map[string]interface{}{
			"Articles": articles[first:count],
			"Site":     globalConfig.Site,
			"Develop":  globalConfig.Develop,
			"Page":     i + 1,
			"Total":    page,
			"Prev":     prev,
			"Next":     next,
			"TagName":  tagName,
			"TagCount": len(articles),
		}
		wg.Add(1)
		go RenderPage(pageTpl, data, outPath)
	}
}

// Generate article list JSON
func GenerateJSON(articles Collections) {
	defer wg.Done()
	datas := make([]map[string]interface{}, 0)
	for i, _ := range articles {
		article := articles[i].(Article)
		var data = map[string]interface{}{
			"title":   article.Title,
			"content": article.Markdown,
			"preview": string(article.Preview),
			"link":    article.Link,
			"cover":   article.Cover,
		}
		datas = append(datas, data)
	}
	str, _ := json.Marshal(datas)
	ioutil.WriteFile(filepath.Join(publicPath, "index.json"), []byte(str), 0644)
}
