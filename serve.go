package main

import (
    "log"
    "os"
    "fmt"
    "net/http"
    "path/filepath"
    "github.com/InkProject/ink.go"
    "github.com/go-fsnotify/fsnotify"
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/gorilla/websocket"
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
							// Fatal(err)
						}
					}
				}
				// case err := <-watcher.Errors:
				// 	Log(err.Error())
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

func Websocket(ctx *ink.Context) {
	// Live reload
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

func Static(watch bool) {
	port := globalConfig.Build.Port
	if port == "" {
		port = "8000"
	}
	web := ink.New()
	if watch {
		web.Get("/live", Websocket)
	}
	// Static
	web.Get("*", ink.Static(filepath.Join(rootPath, "public")))
	// Listen
	Log("Open http://localhost:" + port + "/ to preview")
	web.Listen("0.0.0.0:" + port)
}

func Serve() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)

    router, err := rest.MakeRouter(
        rest.Get("/message", func(w rest.ResponseWriter, req *rest.Request) {
            w.WriteJson(map[string]string{"Body": "Hello World!"})
        }),
        rest.Get("/live", func(w rest.ResponseWriter, req *rest.Request) {
            var upgrader = websocket.Upgrader{
        		ReadBufferSize:  1024,
        		WriteBufferSize: 1024,
        	}
        	if c, err := upgrader.Upgrade(w.(http.ResponseWriter), req.Request, nil); err != nil {
        		Fatal(err)
        	} else {
        		conn = c
        	}
        }),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)

    // http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("editor/assets"))))
    http.Handle("/preview/", http.StripPrefix("/preview", http.FileServer(http.Dir(filepath.Join(rootPath, "public")))))

    log.Fatal(http.ListenAndServe(":8000", nil))
}
