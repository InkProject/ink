package main

import (
    "log"
    "os"
    "fmt"
    "strings"
    "io/ioutil"
    "net/http"
    "path/filepath"
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

func ApiList(w rest.ResponseWriter, req *rest.Request) {
    ret := make([]map[string]string, 0)
    filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
        fileExt := strings.ToLower(filepath.Ext(path))
        if fileExt == ".md" {
            fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
            // Read data from file
            data, err := ioutil.ReadFile(path)
            if err != nil {
                Fatal(err.Error())
            }
            // Split config and markdown
            content := string(data)
            ret = append(ret, map[string]string {
                "name": fileName,
                "content": content,
            })
        }
        return nil
    })
    w.WriteJson(ret)
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
        rest.Get("/list", ApiList),
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

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
