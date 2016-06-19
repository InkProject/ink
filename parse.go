package main

import (
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/InkProject/blackfriday"
	"gopkg.in/yaml.v2"
)

type SiteConfig struct {
	Root     string
	Title    string
	Subtitle string
	Logo     string
	Limit    int
	Theme    string
	Comment  string
	Lang     string
	Url      string
	Link     string
	Config   interface{}
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
	I18n    map[string]string
	Site    SiteConfig
	Authors map[string]AuthorConfig
	Build   BuildConfig
	Develop bool
}

type ArticleConfig struct {
	Title      string
	Date       string
	Update     string
	Author     string
	Tags       []string
	Categories []string
	Topic      string
	Cover      string
	Draft      bool
	Preview    template.HTML
	Top        bool
	Config     interface{}
}

type Article struct {
	GlobalConfig
	ArticleConfig
	Time     time.Time
	MTime    time.Time
	Date     int64
	Update   int64
	Author   AuthorConfig
	Category string
	Tags     []string
	Preview  template.HTML
	Content  template.HTML
	Link     string
	Config   interface{}
}

type ThemeConfig struct {
	Copy []string
	Lang map[string]map[string]string
}

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

func Parse(markdown string) template.HTML {
	// html.UnescapeString
	return template.HTML(blackfriday.MarkdownCommon([]byte(markdown)))
}

func ReplaceRootFlag(content string) string {
	return strings.Replace(content, "-/", globalConfig.Site.Root+"/", -1)
}

func ParseGlobalConfig(configPath string, develop bool) *GlobalConfig {
	var config *GlobalConfig
	// Parse Global Config
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		Fatal(err.Error())
	}
	if config.Site.Config == nil {
		config.Site.Config = ""
	}
	config.Develop = develop
	if develop {
		config.Site.Root = ""
	}
	config.Site.Logo = strings.Replace(config.Site.Logo, "-/", config.Site.Root+"/", -1)
	// Parse Theme Config
	themeConfig := ParseThemeConfig(filepath.Join(rootPath, config.Site.Theme, "config.yml"))
	for _, copyItem := range themeConfig.Copy {
		config.Build.Copy = append(config.Build.Copy, filepath.Join(config.Site.Theme, copyItem))
	}
	config.I18n = make(map[string]string)
	for item, langItem := range themeConfig.Lang {
		config.I18n[item] = langItem[config.Site.Lang]
	}
	return config
}

func ParseThemeConfig(configPath string) *ThemeConfig {
	// Read data from file
	var themeConfig *ThemeConfig
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		Fatal(err.Error())
	}
	// Parse config content
	if err := yaml.Unmarshal(data, &themeConfig); err != nil {
		Fatal(err.Error())
	}
	return themeConfig
}

func ParseArticleConfig(markdownPath string) (config *ArticleConfig, content string) {
	var configStr string
	// Read data from file
	data, err := ioutil.ReadFile(markdownPath)
	if err != nil {
		Fatal(err.Error())
	}
	// Split config and markdown
	contentStr := string(data)
	contentStr = ReplaceRootFlag(contentStr)
	markdownStr := strings.SplitN(contentStr, CONFIG_SPLIT, 2)
	contentLen := len(markdownStr)
	if contentLen > 0 {
		configStr = markdownStr[0]
	}
	if contentLen > 1 {
		content = markdownStr[1]
	}
	// Parse config content
	if err := yaml.Unmarshal([]byte(configStr), &config); err != nil {
		Error(err.Error())
		return nil, ""
	}
	if config == nil {
		return nil, ""
	}
	// Parse preview splited by MORE_SPLIT
	previewAry := strings.SplitN(content, MORE_SPLIT, 2)
	if len(config.Preview) <= 0 && len(previewAry) > 1 {
		config.Preview = Parse(previewAry[0])
		content = strings.Replace(content, MORE_SPLIT, "", 1)
	}
	return config, content
}

func ParseArticle(markdownPath string) *Article {
	config, content := ParseArticleConfig(markdownPath)
	if config == nil {
		Error("Invalid format: " + markdownPath)
		return nil
	}
	if config.Config == nil {
		config.Config = ""
	}
	var article Article
	// Parse markdown content
	article.Preview = config.Preview
	article.Config = config.Config
	article.Content = Parse(content)
	article.Time = ParseDate(config.Date)
	article.Date = article.Time.Unix()
	if config.Update != "" {
		article.MTime = ParseDate(config.Update)
		article.Update = article.MTime.Unix()
	}
	article.Title = config.Title
	article.Topic = config.Topic
	article.Draft = config.Draft
	article.Top = config.Top
	if author, ok := globalConfig.Authors[config.Author]; ok {
		author.Id = config.Author
		author.Avatar = ReplaceRootFlag(author.Avatar)
		article.Author = author
	}
	if len(config.Categories) > 0 {
		article.Category = config.Categories[0]
	} else {
		article.Category = "misc"
	}
	tags := map[string]bool{}
	alltags := append([]string{}, config.Tags...)
	alltags = append(alltags, config.Categories...)
	for _, tag := range alltags {
		tags[tag] = true
	}
	article.Tags = []string{}
	for tag := range tags {
		article.Tags = append(article.Tags, tag)
	}
	// Support topic and cover field
	if config.Cover != "" {
		article.Cover = config.Cover
	} else {
		article.Cover = config.Topic
	}
	return &article
}
