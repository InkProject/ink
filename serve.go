package main

import (
	"net/http"
	"os"
	"path/filepath"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var watcher *fsnotify.Watcher
var conn *websocket.Conn

func buildWatchList() (files []string, dirs []string) {
	dirs = []string{
		filepath.Join(rootPath, "source"),
	}
	files = []string{
		filepath.Join(rootPath, "config.yml"),
		themePath,
	}

	// Add files and directories defined in theme's config.yml to watcher
	for _, themeCopiedPath := range themeConfig.Copy {
		if themeCopiedPath != "" {
			fullPath := filepath.Join(themePath, themeCopiedPath)
			s, err := os.Stat(fullPath)
			if s == nil || err != nil {
				continue
			}

			if s.IsDir() {
				dirs = append(dirs, fullPath)
			} else {
				files = append(files, fullPath)
			}
		}
	}
	return files, dirs
}

// Add files and dirs to watcher
func configureWatcher(watcher *fsnotify.Watcher, files []string, dirs []string) {
	for _, source := range dirs {
		if err := walkSymlinks(source, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				Warn(err.Error())
				return nil
			}
			if f != nil && f.IsDir() {
				if err := watcher.Add(path); err != nil {
					Warn(err.Error())
				}
			}
			return nil
		}); err != nil {
			Warn(err.Error())
		}
	}
	for _, source := range files {
		if err := watcher.Add(source); err != nil {
			Warn(err.Error())
		}
	}
}

func Watch() {
	// Listen watched file change event
	if watcher != nil {
		if err := watcher.Close(); err != nil {
			Warn(err.Error())
		}
	}
	newWatcher, err := fsnotify.NewWatcher()
	if err != nil {
		Fatal(err.Error())
	}
	watcher = newWatcher
	files, dirs := buildWatchList()
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) || event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
					// Handle when file change
					Log(event.Name)
					ParseGlobalConfigWrap(rootPath, true)

					newFiles, newDirs := buildWatchList()
					// If file list changed, reconfigure watcher
					if !reflect.DeepEqual(files, newFiles) || !reflect.DeepEqual(dirs, newDirs) {
						configureWatcher(watcher, newFiles, newDirs)
						files = newFiles
						dirs = newDirs
					}

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
	configureWatcher(watcher, files, dirs)
}

func Websocket(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if c, err := upgrader.Upgrade(w, r, nil); err != nil {
		Warn(err)
	} else {
		conn = c
	}
}

func Serve() {
	// editorWeb := http.NewServeMux()
	//
	// editorWeb.HandleFunc("GET /articles", ApiListArticle)
	// editorWeb.HandleFunc("GET /articles/{id}", ApiGetArticle)
	// editorWeb.HandleFunc("POST /articles", ApiCreateArticle)
	// editorWeb.HandleFunc("PUT /articles/{id}", ApiSaveArticle)
	// editorWeb.HandleFunc("DELETE /articles/{id}", ApiRemoveArticle)
	// editorWeb.HandleFunc("GET /config", ApiGetConfig)
	// editorWeb.HandleFunc("PUT /config", ApiSaveConfig)
	// editorWeb.HandleFunc("POST /upload", ApiUploadFile)
	// editorWeb.Handle("/", http.FileServer(http.Dir(filepath.Join("editor/assets"))))

	// Log("Access http://localhost:" + globalConfig.Build.Port + "/ to open editor")
	// go http.ListenAndServe(":2333", editorWeb)

	previewWeb := http.NewServeMux()
	previewWeb.HandleFunc("/live", Websocket)
	previewWeb.Handle("/", http.FileServer(http.Dir(filepath.Join(rootPath, globalConfig.Build.Output))))

	uri := "http://localhost:" + globalConfig.Build.Port + "/"
	Log("Access " + uri + " to open preview")
	if err := http.ListenAndServe(":"+globalConfig.Build.Port, previewWeb); err != nil {
		Warn(err)
	}
}
