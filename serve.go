package main

import (
	"fmt"
	"github.com/InkProject/ink.go"
	"github.com/go-fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"os"
	"path/filepath"
)

var watcher *fsnotify.Watcher
var conn *websocket.Conn

func Watch() {
	// Listen watched file change event
	if watcher != nil {
		watcher.Close()
	}
	watcher, _ = fsnotify.NewWatcher()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == fsnotify.Write {
					// Handle when file change
					fmt.Println(event.Name)
					Build()
					if conn != nil {
						if err := conn.WriteMessage(websocket.TextMessage, []byte("change")); err != nil {
							Warn(err.Error())
						}
					}
				}
			case err := <-watcher.Errors:
				Warn(err.Error())
			}
		}
	}()
	var dirs = []string{"source"}
	for _, source := range dirs {
		dirPath := filepath.Join(rootPath, source)
		filepath.Walk(dirPath, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				if err := watcher.Add(path); err != nil {
					Warn(err.Error())
				}
			}
			return nil
		})
	}
}

func Websocket(ctx *ink.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if c, err := upgrader.Upgrade(ctx.Res, ctx.Req, nil); err != nil {
		Fatal(err)
	} else {
		conn = c
	}
	ctx.Stop()
}

func Serve() {
	web := ink.New()

	web.Get("/articles", ApiListArticle)
	web.Get("/articles/:id/", ApiGetArticle)
	web.Post("/articles", ApiCreateArticle)
	web.Put("/articles/:id", ApiModifyArticle)
	web.Delete("/articles/:id", ApiRemoveArticle)
	web.Get("/live", Websocket)

	web.Get("*", ink.Static(filepath.Join("editor/assets")))
	// web.Get("*", ink.Static(filepath.Join(rootPath, "public")))
	web.Listen(":" + globalConfig.Build.Port)
}
