package main

import (
    "github.com/codegangsta/cli"
    "github.com/imeoer/bamboo-api/ink"
    "github.com/go-fsnotify/fsnotify"
    "gopkg.in/yaml.v2"
    "os"
    "bufio"
    "strings"
    "runtime"
    "os/exec"
    "io/ioutil"
    "path/filepath"
)

const VERSION = "0.1.0"

var watcher *fsnotify.Watcher
var globalConfig *GlobalConfig
var rootPath string

func main() {
    app := cli.NewApp()
    app.Name = "ink"
    app.Usage = "A concise static blog generator"
    app.Author = "https://github.com/imeoer"
    app.Email = "imeoer@gmail.com"
    app.Version = VERSION
    app.Commands = []cli.Command{
        {
            Name: "preview",
            Usage: "Run in server mode to preview blog",
            Action: func(c *cli.Context) {
                ParseGlobalConfig(c)
                globalConfig.Develop = true
                Build()
                Watch()
                Server()
            },
        },
        {
            Name: "publish",
            Usage: "Generate blog to public folder and publish",
            Action: func(c *cli.Context) {
                ParseGlobalConfig(c)
                Build()
                Publish()
            },
        },
        {
            Name: "convert",
            Usage: "Convert Jekyll/Hexo post format to Ink format",
            Action: func(c *cli.Context) {
                Convert(c)
            },
        },
    }
    app.Action = func(c *cli.Context) {
        ParseGlobalConfig(c)
        Build()
    }
    app.Run(os.Args)
}

func ParseGlobalConfig(c *cli.Context) {
    if len(c.Args()) > 0 {
        rootPath = c.Args()[0]
    } else {
        rootPath = "."
    }
    globalConfig = ParseConfig(filepath.Join(rootPath, "config.yml"))
}

func Server() {
    port := globalConfig.Build.Port
    if port == "" {
        port = "8888"
    }
    app := ink.New()
    app.Get("*", ink.Static(rootPath+"/public"))
    app.Head("*", ink.Static(rootPath+"/public"))
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
        filepath.Walk(dirPath, func (path string, f os.FileInfo, err error) error {
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
    filepath.Walk(sourcePath, func (path string, f os.FileInfo, err error) error {
        fileExt := strings.ToLower(filepath.Ext(path))
        if fileExt == ".md" || fileExt == ".html" {
            // Read data from file
            data, err := ioutil.ReadFile(path)
            if err != nil {
                Fatal(err.Error())
            }
            // Split config and markdown
            var configStr, contentStr string
            content := strings.TrimSpace(string(data))
            parseAry := strings.SplitN(content, "---", 3)
            parseLen := len(parseAry)
            if parseLen == 3 { // jekyll
                configStr = parseAry[1]
                contentStr = parseAry[2]
            } else if parseLen == 2 { // hexo
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
            dateAry := strings.SplitN(article.Date, ".", 2)
            if len(dateAry) == 2 {
                article.Date = dateAry[0]
            }
            // Generate Config
            var inkConfig []byte
            if inkConfig, err = yaml.Marshal(article); err != nil {
                Fatal(err.Error())
            }
            inkConfigStr := string(inkConfig)
            markdownStr := inkConfigStr + "\n\n---\n\n" + contentStr
            fileName := filepath.Base(path)
            ioutil.WriteFile(filepath.Join(rootPath, "source/" + fileName + ".md"), []byte(markdownStr), 0666)
        }
        return nil
    })
}
