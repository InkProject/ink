package main

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/go-fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
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
						if err := conn.WriteMessage(websocket.TextMessage, []byte("change")); err == nil {
							Log(err)
						}
					}
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
				if err := watcher.Add(path); err != nil {
					Log(err.Error())
				}
			}
			return nil
		})
	}
}

func Websocket(w rest.ResponseWriter, req *rest.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if c, err := upgrader.Upgrade(w.(http.ResponseWriter), req.Request, nil); err != nil {
		Fatal(err)
	} else {
		conn = c
	}
}

func Serve() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/articles", ApiListArticle),
		rest.Get("/articles/#id", ApiGetArticle),
		rest.Post("/articles", ApiCreateArticle),
		rest.Put("/articles/#id", ApiModifyArticle),
		rest.Delete("/articles/#id", ApiRemoveArticle),
		rest.Get("/live", Websocket),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)

	http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
	http.Handle("/editor/", http.StripPrefix("/editor", http.FileServer(http.Dir("editor/assets"))))
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(filepath.Join(rootPath, "public")))))

	port := globalConfig.Build.Port
	if port == "" {
		port = "8000"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
