package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"github.com/InkProject/ink.go"
	"github.com/codegangsta/cli"
	"github.com/go-fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	VERSION      = "Beta (2015-06-04)"
	DEFAULT_PATH = "blog"
	DOWNLOAD_URL = "http://www.inkpaper.io/release/ink_blog.zip"
)

var watcher *fsnotify.Watcher
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
		// {
		// 	Name:  "init",
		// 	Usage: "Init blog in a specified directory",
		// 	Action: func(c *cli.Context) {
		// 		Init(c)
		// 	},
		// },
		{
			Name:  "preview",
			Usage: "Run in server mode to preview blog",
			Action: func(c *cli.Context) {
				ParseGlobalConfigByCli(c, true)
				Build()
				Watch()
				Server()
			},
		},
		{
			Name:  "publish",
			Usage: "Generate blog to public folder and publish",
			Action: func(c *cli.Context) {
				ParseGlobalConfigByCli(c, false)
				Build()
				Publish()
			},
		},
		{
			Name:  "convert",
			Usage: "Convert Jekyll/Hexo post format to Ink format",
			Action: func(c *cli.Context) {
				Convert(c)
			},
		},
	}
	app.Action = func(c *cli.Context) {
		ParseGlobalConfig(".", false)
		if globalConfig == nil {
			ParseGlobalConfig(DEFAULT_PATH, false)
			if globalConfig == nil {
				// Init(nil)
				// ParseGlobalConfig(DEFAULT_PATH, true)
				Fatal("Config.yml not found, please specify a valid blog directory")
			}
		}
		if globalConfig != nil {
			Build()
			Watch()
			Server()
		}
	}
	app.Run(os.Args)
}

func ParseGlobalConfigByCli(c *cli.Context, develop bool) {
	if len(c.Args()) > 0 {
		rootPath = c.Args()[0]
	} else {
		rootPath = "."
	}
	ParseGlobalConfig(rootPath, develop)
	if globalConfig == nil {
		Fatal("Parse config.yml failed, please specify a valid path")
	}
}

func ParseGlobalConfig(root string, develop bool) {
	rootPath = root
	globalConfig = ParseConfig(filepath.Join(rootPath, "config.yml"))
	if globalConfig == nil {
		return
	}
	globalConfig.Develop = develop
	if develop {
		globalConfig.Site.Root = ""
	}
	globalConfig.Site.Logo = ReplaceRootFlag(globalConfig.Site.Logo)
}

func Server() {
	port := globalConfig.Build.Port
	if port == "" {
		port = "8888"
	}
	app := ink.New()
	app.Get("*", ink.Static(filepath.Join(rootPath, "public")))
	app.Head("*", ink.Static(filepath.Join(rootPath, "public")))
	Log("Open http://localhost:" + port + "/ to preview")
	app.Listen("0.0.0.0:" + port)
}

func Watch() {
	watcher, _ = fsnotify.NewWatcher()
	// Listen watched file change event
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Write {
					// Handle when file change
					Build()
				}
			case err := <-watcher.Errors:
				Log(err.Error())
			}
		}
	}()
	var dirs = []string{"source"}
	for _, source := range dirs {
		dirPath := filepath.Join(rootPath, source)
		filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				// Defer watcher.Close()
				if err := watcher.Add(path); err != nil {
					Log(err.Error())
				}
			}
			return nil
		})
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
	cmd.Dir = filepath.Join(rootPath, "public")
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
	filepath.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
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
			markdownStr := inkConfigStr + "\n\n---\n\n" + contentStr
			ioutil.WriteFile(filepath.Join(rootPath, "source/"+fileName+".md"), []byte(markdownStr), 0666)
			count++
		}
		return nil
	})
	fmt.Printf("\nConvert finish, total %v articles\n", count)
}

// Download and extract Ink template

type Download struct {
	io.Reader
	length int64
	total  int64
}

func (dn *Download) Read(p []byte) (int, error) {
	n, err := dn.Reader.Read(p)
	dn.total += int64(n)
	if err == nil {
		percent := math.Ceil(float64(dn.total) / float64(dn.length) * 100)
		fmt.Printf("\rDownload %.f%% ...\r", percent)
	}
	return n, err
}

func Init(c *cli.Context) {
	// Parse arguments
	var directory string
	if c == nil {
		directory = DEFAULT_PATH
	} else {
		args := c.Args()
		if len(args) > 0 {
			directory = args[0]
		} else {
			Fatal("Please specify a new blog directory name")
		}
	}
	// Create blog directory
	err := os.MkdirAll(directory, 0777)
	if err != nil {
		Fatal(err.Error())
	}
	zipPath := filepath.Join(os.TempDir(), "ink_blog.zip")
	zipOut, err := os.Create(zipPath)
	if err != nil {
		Fatal(err.Error())
	}
	defer zipOut.Close()
	// Http get request for zip
	fmt.Printf("Connecting server to init blog\r")
	resp, err := http.Get(DOWNLOAD_URL)
	if err != nil {
		Fatal(err.Error())
	}
	defer resp.Body.Close()
	// Get download progress
	zipSrc := &Download{Reader: resp.Body, length: resp.ContentLength}
	_, err = io.Copy(zipOut, zipSrc)
	if err != nil {
		Fatal(err.Error())
	}
	// Extract downloaded zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		Fatal(err.Error())
	}
	defer r.Close()
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			Fatal(err.Error())
		}
		isDir := f.FileInfo().IsDir()
		if isDir {
			err = os.MkdirAll(filepath.Join(directory, f.Name), 0777)
			if err != nil {
				Fatal(err.Error())
			}
		} else {
			extractOut, err := os.Create(filepath.Join(directory, f.Name))
			if err != nil {
				Fatal(err.Error())
			}
			_, err = io.Copy(extractOut, rc)
			if err != nil {
				Fatal(err.Error())
			}
		}
		rc.Close()
	}
	Log("Blog created in " + directory + ", use 'ink preview " + directory + "' to preview")
}
