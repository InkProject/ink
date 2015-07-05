package main

import (
	"github.com/InkProject/blackfriday"
	"gopkg.in/yaml.v2"
	"html/template"
	"io/ioutil"
	"strings"
	"path/filepath"
)

type SiteConfig struct {
	Root     string
	Title    string
	Subtitle string
	Logo     string
	Limit    int
	Theme    string
	Disqus   string
	Lang  	string
}

type AuthorConfig struct {
	Id     string
	Name   string
	Intro  string
	Avatar string
}

type BuildConfig struct {
	Port    string
	Watch   bool
	Copy    []string
	Publish string
}

type GlobalConfig struct {
	I18n	map[string]string
	Site    SiteConfig
	Authors map[string]AuthorConfig
	Build   BuildConfig
	Develop bool
}

type ArticleConfig struct {
	Title   string
	Date    string
	Update  string
	Author  string
	Tags    []string
	Topic   string
	Draft   bool
	Preview string
	Top     bool
}

type Article struct {
	ArticleConfig
	GlobalConfig
	Date    int64
	Update  int64
	Author  AuthorConfig
	Tags    []string
	Preview string
	Content template.HTML
	Link    string
}

type Lang interface{}

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

func parse(markdown string) template.HTML {
	// html.UnescapeString
	return template.HTML(blackfriday.MarkdownCommon([]byte(markdown)))
}

func ReplaceRootFlag(content string) string {
	return strings.Replace(content, "-/", globalConfig.Site.Root+"/", -1)
}

func ParseConfig(configPath string, develop bool) *GlobalConfig {
	var config *GlobalConfig
	// Read data from file
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		Fatal(err.Error())
	}
	lang := ParseLang(filepath.Join(rootPath, config.Site.Theme, "lang.yml"))
	config.I18n = make(map[string]string)
	for item, langItem := range lang {
		config.I18n[item] = langItem[config.Site.Lang]
	}
	config.Develop = develop
	if develop {
		config.Site.Root = ""
	}
	config.Site.Logo = strings.Replace(config.Site.Logo, "-/", config.Site.Root+"/", -1)
	return config
}

func ParseLang(langPath string) map[string]map[string]string {
	// Read data from file
	var lang map[string]map[string]string
	data, err := ioutil.ReadFile(langPath)
	if err != nil {
		Fatal(err.Error())
	}
	// Parse lang content
	if err := yaml.Unmarshal(data, &lang); err != nil {
		Fatal(err.Error())
	}
	return lang
}

func ParseMarkdown(markdownPath string) *Article {
	var (
		config      *ArticleConfig
		configStr   string
		markdownStr string
	)
	// Read data from file
	data, err := ioutil.ReadFile(markdownPath)
	if err != nil {
		Fatal(err.Error())
	}
	// Split config and markdown
	contentStr := string(data)
	contentStr = ReplaceRootFlag(contentStr)
	content := strings.SplitN(contentStr, CONFIG_SPLIT, 2)
	contentLen := len(content)
	if contentLen > 0 {
		configStr = content[0]
	}
	if contentLen > 1 {
		markdownStr = content[1]
	}
	// Parse config content
	if err := yaml.Unmarshal([]byte(configStr), &config); err != nil {
		Fatal(err.Error())
	}
	if config == nil {
		Fatal("Article config parse error")
	}
	var article Article
	// Parse preview splited by MORE_SPLIT
	previewAry := strings.SplitN(markdownStr, MORE_SPLIT, 2)
	if len(config.Preview) > 0 {
		article.Preview = config.Preview
	} else {
		if len(previewAry) > 1 {
			article.Preview = previewAry[0]
			markdownStr = strings.Replace(markdownStr, MORE_SPLIT, "", 1)
		}
	}
	// Parse markdown content
	article.Content = parse(markdownStr)
	article.Date = ParseDate(config.Date).Unix()
	if config.Update != "" {
		article.Update = ParseDate(config.Update).Unix()
	}
	article.Title = config.Title
	if author, ok := globalConfig.Authors[config.Author]; ok {
		author.Id = config.Author
		author.Avatar = ReplaceRootFlag(author.Avatar)
		article.Author = author
	}
	article.Tags = config.Tags
	article.Topic = config.Topic
	article.Draft = config.Draft
	article.Top = config.Top
	return &article
}
