package main

import (
	"os"
	"path/filepath"

	"github.com/InkProject/ink.go"
	"github.com/facebookgo/symwalk"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/taadis/oper"
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
					Log(event.Name)
					ParseGlobalConfigWrap(rootPath, true)
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
	var dirs = []string{
		filepath.Join(rootPath, "source"),
		filepath.Join(themePath, "bundle"),
	}
	var files = []string{
		filepath.Join(rootPath, "config.yml"),
		filepath.Join(themePath),
	}
	for _, source := range dirs {
		symwalk.Walk(source, func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				if err := watcher.Add(path); err != nil {
					Warn(err.Error())
				}
			}
			return nil
		})
	}
	for _, source := range files {
		if err := watcher.Add(source); err != nil {
			Warn(err.Error())
		}
	}
}

func Websocket(ctx *ink.Context) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if c, err := upgrader.Upgrade(ctx.Res, ctx.Req, nil); err != nil {
		Warn(err)
	} else {
		conn = c
	}
	ctx.Stop()
}

func Serve() {
	// editorWeb := ink.New()
	//
	// editorWeb.Get("/articles", ApiListArticle)
	// editorWeb.Get("/articles/:id", ApiGetArticle)
	// editorWeb.Post("/articles", ApiCreateArticle)
	// editorWeb.Put("/articles/:id", ApiSaveArticle)
	// editorWeb.Delete("/articles/:id", ApiRemoveArticle)
	// editorWeb.Get("/config", ApiGetConfig)
	// editorWeb.Put("/config", ApiSaveConfig)
	// editorWeb.Post("/upload", ApiUploadFile)
	// editorWeb.Use(ink.Cors)
	// editorWeb.Get("*", ink.Static(filepath.Join("editor/assets")))

	// Log("Access http://localhost:" + globalConfig.Build.Port + "/ to open editor")
	// go editorWeb.Listen(":2333")

	previewWeb := ink.New()
	previewWeb.Get("/live", Websocket)
	previewWeb.Get("*", ink.Static(filepath.Join(rootPath, globalConfig.Build.Output)))

	uri := "http://localhost:" + globalConfig.Build.Port + "/"
	Log("Access " + uri + " to open preview")
	oper.Access(uri)
	previewWeb.Listen(":" + globalConfig.Build.Port)
}
