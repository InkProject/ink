package main

import (
    "github.com/codegangsta/cli"
    "github.com/imeoer/bamboo-api/ink"
    "os"
    "bufio"
    "runtime"
    "os/exec"
    "path/filepath"
)

var globalConfig *GlobalConfig
var rootPath string

func main() {
    app := cli.NewApp()
    app.Name = "Ink"
    app.Usage = "A static blog generator"
    app.Author = "https://github.com/imeoer"
    app.Email = "imeoer@gmail.com"
    app.Version = "0.1.0"
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "server, s",
            Usage: "Run in server mode to preview site",
        },
    }
    app.Commands = []cli.Command{
        {
            Name: "publish",
            ShortName: "p",
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
        serverMode := c.Bool("server")
        if serverMode {
            Server()
        }
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
    Log(LOG, "Server running on port "+port)
    app.Listen("0.0.0.0:" + port)
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
