package main

import (
    "os"
    "github.com/imeoer/bamboo-api/ink"
    "github.com/codegangsta/cli"
)

func main() {
    var root string
    app := cli.NewApp()
    app.Name = "Ink"
    app.Usage = "A static blog generator"
    app.Author = "https://github.com/imeoer"
    app.Email = "imeoer@gmail.com"
    app.Version = "0.1.0"
    app.Flags = []cli.Flag{
        cli.BoolFlag{
            Name: "server, s",
            Usage: "Run a server to preview site",
        },
    }
    app.Action = func(c *cli.Context) {
        if len(c.Args()) > 0 {
            root = c.Args()[0]
        } else {
            root = "."
        }
        Create(root)
        serverMode := c.Bool("server")
        if serverMode {
            app := ink.New()
            app.Get("*", ink.Static(root + "/public"))
            Log(LOG, "Server running on port 8888")
            app.Listen("0.0.0.0:8888")
        }
    }
    app.Run(os.Args)
}
