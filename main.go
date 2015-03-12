package main

import (
    "github.com/codegangsta/cli"
    "github.com/imeoer/bamboo-api/ink"
    "github.com/go-fsnotify/fsnotify"
    "os"
    "bufio"
    "runtime"
    "os/exec"
    "path/filepath"
)

var watcher *fsnotify.Watcher
var globalConfig *GlobalConfig
var rootPath string

func main() {
    app := cli.NewApp()
    app.Name = "Ink"
    app.Usage = "A static blog generator"
    app.Author = "https://github.com/imeoer"
    app.Email = "imeoer@gmail.com"
    app.Version = "0.1.0"
    app.Commands = []cli.Command{
        {
            Name: "preview",
            ShortName: "pre",
            Usage: "Run in server mode to preview site",
            Action: func(c *cli.Context) {
                ParseGlobalConfig(c)
                Build()
                Watch()
                Server()
            },
        },
        {
            Name: "publish",
            ShortName: "pub",
            Usage: "Publish all files in public folder",
            Action: func(c *cli.Context) {
                ParseGlobalConfig(c)
                Build()
                Publish()
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
    Log(LOG, "Listening on port "+port)
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
                    Log(ERR, err.Error())
            }
        }
    }()
    var dirs = []string{"theme", "source"}
    for _, source := range dirs {
        dirPath := filepath.Join(rootPath, source)
        filepath.Walk(dirPath, func (path string, f os.FileInfo, err error) error {
            if f.IsDir() {
                // Defer watcher.Close()
                if err := watcher.Add(path); err != nil {
                    Log(ERR, err.Error())
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
            Log(LOG, out.Text())
        }
    }()
    // Print stdin
    go func() {
        for err.Scan() {
            Log(LOG, err.Text())
        }
    }()
    // Exec command
    cmd.Run()
}
