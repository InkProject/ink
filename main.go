package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/facebookgo/symwalk"
	"gopkg.in/yaml.v2"
	"html/template"
	"time"
)

const (
	VERSION            = "RELEASE 2018-07-27"
	DEFAULT_ROOT       = "blog"
	DATE_FORMAT_STRING = "2006-01-02 15:04:05"
	INDENT             = "  " // 2 spaces
	POST_TEMPLATE      = `title: {{.Title}}
date: {{.DateString}}
author: {{.Author}}
{{- if .Cover}}
cover: {{.Cover}}
{{- end}}
draft: {{.Draft}}
top: {{.Top}}
{{- if .Preview}}
preview: {{.Preview}}
{{- end}}
{{- if .Tags}}
{{.Tags}}
{{- end}}
type: {{.Type}}
hide: {{.Hide}}
---
`
)

var globalConfig *GlobalConfig
var rootPath string

func main() {
	app := cli.NewApp()
	app.Name = "ink"
	app.Usage = "An elegant static blog generator"
	app.Author = "https://github.com/imeoer"
	app.Email = "imeoer@gmail.com"
	app.Version = VERSION
	app.Commands = []cli.Command{
		{
			Name:  "build",
			Usage: "Generate blog to public folder",
			Action: func(c *cli.Context) error {
				ParseGlobalConfigByCli(c, false)
				Build()
				return nil
			},
		},
		{
			Name:  "preview",
			Usage: "Run in server mode to preview blog",
			Action: func(c *cli.Context) error {
				ParseGlobalConfigByCli(c, true)
				Build()
				Watch()
				Serve()
				return nil
			},
		},
		{
			Name:  "publish",
			Usage: "Generate blog to public folder and publish",
			Action: func(c *cli.Context) error {
				ParseGlobalConfigByCli(c, false)
				Build()
				Publish()
				return nil
			},
		},
		{
			Name:  "serve",
			Usage: "Run in server mode to serve blog",
			Action: func(c *cli.Context) error {
				ParseGlobalConfigByCli(c, true)
				Build()
				Serve()
				return nil
			},
		},
		{
			Name:  "convert",
			Usage: "Convert Jekyll/Hexo post format to Ink format (Beta)",
			Action: func(c *cli.Context) error {
				Convert(c)
				return nil
			},
		},
		{
			Name:  "new",
			Usage: "Creates a new article",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "hide",
					Usage: "Hides the article",
				},
				cli.BoolFlag{
					Name:  "top",
					Usage: "Places the article at the top",
				},
				cli.BoolFlag{
					Name:  "post",
					Usage: "The article is a post",
				},
				cli.BoolFlag{
					Name:  "page",
					Usage: "The article is a page",
				},
				cli.BoolFlag{
					Name:  "draft",
					Usage: "The article is a draft",
				},

				cli.StringFlag{
					Name:  "title",
					Usage: "Article title",
				},
				cli.StringFlag{
					Name:  "author",
					Usage: "Article author",
				},
				cli.StringFlag{
					Name:  "cover",
					Usage: "Article cover path",
				},
				cli.StringFlag{
					Name:  "date",
					Usage: "The date and time on which the article was created (2006-01-02 15:04:05)",
				},
				cli.StringFlag{
					Name:  "file",
					Usage: "The path of where the article will be stored",
				},

				cli.StringSliceFlag{
					Name:  "tag",
					Usage: "Adds a tag to the article",
				},
			},
			Action: func(c *cli.Context) error {
				New(c)
				return nil
			},
		},
	}
	app.Run(os.Args)
	os.Exit(exitCode)
}

func ParseGlobalConfigByCli(c *cli.Context, develop bool) {
	if len(c.Args()) > 0 {
		rootPath = c.Args()[0]
	} else {
		rootPath = "."
	}
	ParseGlobalConfigWrap(rootPath, develop)
	if globalConfig == nil {
		ParseGlobalConfigWrap(DEFAULT_ROOT, develop)
		if globalConfig == nil {
			Fatal("Parse config.yml failed, please specify a valid path")
		}
	}
}

func ParseGlobalConfigWrap(root string, develop bool) {
	rootPath = root
	globalConfig = ParseGlobalConfig(filepath.Join(rootPath, "config.yml"), develop)
	if globalConfig == nil {
		return
	}
}

func New(c *cli.Context) {
	// If source folder does not exist, create
	if _, err := os.Stat("source/"); os.IsNotExist(err) {
		os.Mkdir("source", os.ModePerm)
	}

	var author, blogTitle, fileName string
	var tags []string

	// Default values
	draft := "false"
	top := "false"
	postType := "post"
	hide := "false"
	date := time.Now()

	// Empty string values
	preview := ""
	cover := ""

	// Parse args
	args := c.Args()
	if len(args) > 0 {
		blogTitle = args[0]
	}
	if blogTitle == "" {
		if c.String("title") != "" {
			blogTitle = c.String("title")
		} else {
			Fatal("Please specify the name of the blog post")
		}
	}

	fileName = blogTitle + ".md"
	if c.String("file") != "" {
		fileName = c.String("file")
	}

	if len(args) > 1 {
		author = args[1]
	}
	if author == "" {
		author = c.String("author")
	}

	if c.Bool("post") && c.Bool("page") {
		Fatal("The post and page arguments are mutually exclusive and cannot appear together")
	}
	if c.Bool("post") {
		postType = "post"
	}
	if c.Bool("page") {
		postType = "page"
	}
	if c.Bool("hide") {
		hide = "true"
	}
	if c.Bool("draft") {
		draft = "true"
	}
	if c.Bool("top") {
		top = "true"
	}

	if c.String("preview") != "" {
		preview = c.String("preview")
	}
	if c.String("cover") != "" {
		cover = c.String("cover")
	}

	var filePath = "source/" + fileName
	file, err := os.Create(filePath)
	if err != nil {
		Fatal(err)
	}
	postTemplate, err := template.New("post").Parse(POST_TEMPLATE)
	if err != nil {
		Fatal(err)
	}

	if c.StringSlice("tag") != nil {
		tags = c.StringSlice("tag")
	}

	var tagString string
	if len(tags) > 0 {
		tagString = "tags:"
		for _, tag := range tags {
			tagString += "\n" + INDENT + "- " + tag
		}
	}

	var dateString string
	if c.String("date") != "" {
		dateString = c.String("date")
		_, err = time.Parse(DATE_FORMAT_STRING, dateString)
		if err != nil {
			Fatal("Illegal date string")
		}
	} else {
		dateString = date.Format(DATE_FORMAT_STRING)
	}
	data := map[string]string{
		"Title":      blogTitle,
		"DateString": dateString,
		"Author":     author,
		"Draft":      draft,
		"Top":        top,
		"Type":       postType,
		"Hide":       hide,
		"Preview":    preview,
		"Cover":      cover,
		"Tags":       tagString,
	}
	fileWriter := bufio.NewWriter(file)
	err = postTemplate.Execute(fileWriter, data)
	if err != nil {
		Fatal(err)
	}
	err = fileWriter.Flush()
	if err != nil {
		Fatal(err)
	}
}

func Publish() {
	command := globalConfig.Build.Publish
	// Prepare exec command
	var shell, flag string
	if runtime.GOOS == "windows" {
		shell = "cmd"
		flag = "/C"
	} else {
		shell = "/bin/sh"
		flag = "-c"
	}
	cmd := exec.Command(shell, flag, command)
	cmd.Dir = filepath.Join(rootPath, globalConfig.Build.Output)
	// Start print stdout and stderr of process
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	out := bufio.NewScanner(stdout)
	err := bufio.NewScanner(stderr)
	// Print stdout
	go func() {
		for out.Scan() {
			Log(out.Text())
		}
	}()
	// Print stdin
	go func() {
		for err.Scan() {
			Log(err.Text())
		}
	}()
	// Exec command
	cmd.Run()
}

func Convert(c *cli.Context) {
	// Parse arguments
	var sourcePath, rootPath string
	args := c.Args()
	if len(args) > 0 {
		sourcePath = args[0]
	} else {
		Fatal("Please specify the posts source path")
	}
	if len(args) > 1 {
		rootPath = args[1]
	} else {
		rootPath = "."
	}
	// Check if path exist
	if !Exists(sourcePath) || !Exists(rootPath) {
		Fatal("Please specify valid path")
	}
	// Parse Jekyll/Hexo post file
	count := 0
	symwalk.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
		fileExt := strings.ToLower(filepath.Ext(path))
		if fileExt == ".md" || fileExt == ".html" {
			// Read data from file
			data, err := ioutil.ReadFile(path)
			fileName := filepath.Base(path)
			Log("Converting " + fileName)
			if err != nil {
				Fatal(err.Error())
			}
			// Split config and markdown
			var configStr, contentStr string
			content := strings.TrimSpace(string(data))
			parseAry := strings.SplitN(content, "---", 3)
			parseLen := len(parseAry)
			if parseLen == 3 { // Jekyll
				configStr = parseAry[1]
				contentStr = parseAry[2]
			} else if parseLen == 2 { // Hexo
				configStr = parseAry[0]
				contentStr = parseAry[1]
			}
			// Parse config
			var article ArticleConfig
			if err = yaml.Unmarshal([]byte(configStr), &article); err != nil {
				Fatal(err.Error())
			}
			tags := make(map[string]bool)
			for _, t := range article.Tags {
				tags[t] = true
			}
			for _, c := range article.Categories {
				if _, ok := tags[c]; !ok {
					article.Tags = append(article.Tags, c)
				}
			}
			if article.Author == "" {
				article.Author = "me"
			}
			// Convert date
			dateAry := strings.SplitN(article.Date, ".", 2)
			if len(dateAry) == 2 {
				article.Date = dateAry[0]
			}
			if len(article.Date) == 10 {
				article.Date = article.Date + " 00:00:00"
			}
			if len(article.Date) == 0 {
				article.Date = "1970-01-01 00:00:00"
			}
			article.Update = ""
			// Generate Config
			var inkConfig []byte
			if inkConfig, err = yaml.Marshal(article); err != nil {
				Fatal(err.Error())
			}
			inkConfigStr := string(inkConfig)
			markdownStr := inkConfigStr + "\n\n---\n\n" + contentStr + "\n"
			targetName := "source/" + fileName
			if fileExt != ".md" {
				targetName = targetName + ".md"
			}
			ioutil.WriteFile(filepath.Join(rootPath, targetName), []byte(markdownStr), 0644)
			count++
		}
		return nil
	})
	fmt.Printf("\nConvert finish, total %v articles\n", count)
}
