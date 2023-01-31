package main

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	gomk "github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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
	Output  string
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
	Type       string
	Hide       bool
	Toc        bool
	Image      string
	Subtitle   string
	Config     map[string]interface{}
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
	Markdown string
	Preview  template.HTML
	Content  template.HTML
	Link     string
	Config   interface{}
	Image    string
	Subtitle string
}

type ThemeConfig struct {
	Copy []string
	Lang map[string]map[string]string
}

const (
	CONFIG_SPLIT = "---"
	MORE_SPLIT   = "<!--more-->"
)

// Modify the rendering of the image node to use jQuery-unveil to lazy load images.
//
// Note: This is a hacky way to do this, but it works for now.
// LazyLoadImages in markdown.html.Flags should not be used with this hook, and
// adding the absolute path prefix to the image destination (which is what the original parser does)
// is not implemented yet.
func renderHookLazyLoadImage(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	// Skip all nodes that are not Image nodes
	if _, ok := node.(*ast.Image); !ok {
		return ast.GoToNext, false
	}
	img := node.(*ast.Image)
	if entering {
		if img.Attribute == nil {
			img.Attribute = &ast.Attribute{}
		}
		if img.Attrs == nil {
			img.Attrs = make(map[string][]byte)
		}
		img.Attrs["data-src"] = img.Destination
		img.Destination = []byte("data:image/gif;base64,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")
	} else {
		w.Write([]byte("\" data-src=\""))
		w.Write(img.Attrs["data-src"])
	}
	return ast.GoToNext, false
}

func ParseMarkdown(markdown string, toc bool) template.HTML {
	extensions := parser.CommonExtensions | parser.Footnotes
	parser := parser.NewWithExtensions(extensions)

	htmlFlags := html.CommonFlags
	if toc {
		htmlFlags |= html.TOC
	}
	opts := html.RendererOptions{Flags: htmlFlags, RenderNodeHook: renderHookLazyLoadImage}
	renderer := html.NewRenderer(opts)

	return template.HTML(gomk.ToHTML(gomk.NormalizeNewlines([]byte(markdown)), parser, renderer))
}

func ReplaceRootFlag(content string) string {
	return strings.Replace(content, "-/", globalConfig.Site.Root+"/", -1)
}

func ParseGlobalConfig(configPath string, develop bool) *GlobalConfig {
	var config *GlobalConfig
	// Parse Global Config
	data, err := os.ReadFile(configPath)
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
	if config.Site.Url != "" && strings.HasSuffix(config.Site.Url, "/") {
		config.Site.Url = strings.TrimSuffix(config.Site.Url, "/")
	}
	if config.Build.Output == "" {
		config.Build.Output = "public"
	}
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
	data, err := os.ReadFile(configPath)
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
	data, err := os.ReadFile(markdownPath)
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
	if config.Type == "" {
		config.Type = "post"
	}
	// Parse preview splited by MORE_SPLIT
	previewAry := strings.SplitN(content, MORE_SPLIT, 2)
	if len(config.Preview) <= 0 && len(previewAry) > 1 {
		config.Preview = ParseMarkdown(previewAry[0], false)
		content = strings.Replace(content, MORE_SPLIT, "", 1)
	} else {
		config.Preview = ParseMarkdown(string(config.Preview), false)
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
		config.Config = make(map[string]interface{})
	}
	var article Article
	// Parse markdown content
	article.Hide = config.Hide
	article.Type = config.Type
	article.Preview = config.Preview
	article.Config = config.Config
	article.Markdown = content
	article.Content = ParseMarkdown(content, config.Toc)
	if config.Date != "" {
		article.Time = ParseDate(config.Date)
		article.Date = article.Time.Unix()
	}
	if config.Update != "" {
		article.MTime = ParseDate(config.Update)
		article.Update = article.MTime.Unix()
	}
	article.Title = config.Title
	article.Topic = config.Topic
	article.Draft = config.Draft
	article.Top = config.Top
	article.Image = config.Image
	article.Subtitle = config.Subtitle
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
	article.Tags = config.Tags
	for _, tag := range config.Tags {
		tags[tag] = true
	}
	for _, cat := range config.Categories {
		if _, ok := tags[cat]; !ok {
			article.Tags = append(article.Tags, cat)
		}
	}
	// Support topic and cover field
	if config.Cover != "" {
		article.Cover = config.Cover
	} else {
		article.Cover = config.Topic
	}
	// Generate page name
	fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(markdownPath)), ".md")
	link := fileName + ".html"
	// Genetate custom link
	if article.Type == "post" {
		datePrefix := article.Time.Format("2006-01-02-")
		fileName = strings.TrimPrefix(fileName, datePrefix)
		if globalConfig.Site.Link != "" {
			linkMap := map[string]string{
				"{year}":     article.Time.Format("2006"),
				"{month}":    article.Time.Format("01"),
				"{day}":      article.Time.Format("02"),
				"{hour}":     article.Time.Format("15"),
				"{minute}":   article.Time.Format("04"),
				"{second}":   article.Time.Format("05"),
				"{category}": article.Category,
				"{title}":    fileName,
			}
			link = globalConfig.Site.Link
			for key, val := range linkMap {
				link = strings.Replace(link, key, val, -1)
			}
		}
	}
	article.Link = link
	article.GlobalConfig = *globalConfig
	return &article
}
