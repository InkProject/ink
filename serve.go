package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var watcher *fsnotify.Watcher
var reloadMu sync.Mutex
var reloadClients = make(map[*websocket.Conn]struct{})

const watchDebounce = 150 * time.Millisecond
const reloadWriteTimeout = 5 * time.Second

func notifyReloadClients() {
	reloadMu.Lock()
	defer reloadMu.Unlock()
	for client := range reloadClients {
		err := client.SetWriteDeadline(time.Now().Add(reloadWriteTimeout))
		if err == nil {
			err = client.WriteMessage(websocket.TextMessage, []byte("change"))
		}
		if err == nil {
			continue
		}
		Warn(err.Error())
		if err := client.Close(); err != nil {
			Warn(err.Error())
		}
		delete(reloadClients, client)
	}
}

func buildWatchList() (files []string, dirs []string) {
	configuredThemePath := filepath.Join(rootPath, globalConfig.Site.Theme)
	dirs = []string{
		filepath.Join(rootPath, "source"),
	}
	files = []string{
		filepath.Join(rootPath, "config.yml"),
		configuredThemePath,
	}

	// Add files and directories defined in theme's config.yml to watcher
	for _, themeCopiedPath := range themeConfig.Copy {
		if themeCopiedPath != "" {
			fullPath := filepath.Join(configuredThemePath, themeCopiedPath)
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

// Make the active watch set exactly match the current configuration.
func configureWatcher(watcher *fsnotify.Watcher) error {
	files, dirs := buildWatchList()
	desired := make(map[string]struct{})
	for _, source := range dirs {
		if err := walkSymlinks(source, func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if f != nil && f.IsDir() {
				desired[filepath.Clean(path)] = struct{}{}
			}
			return nil
		}); err != nil {
			return err
		}
	}
	for _, source := range files {
		desired[filepath.Clean(source)] = struct{}{}
	}
	for _, path := range watcher.WatchList() {
		path = filepath.Clean(path)
		if _, ok := desired[path]; ok {
			delete(desired, path)
			continue
		}
		if err := watcher.Remove(path); err != nil && !errors.Is(err, fsnotify.ErrNonExistentWatch) {
			return fmt.Errorf("remove watch %q: %w", path, err)
		}
	}
	for path := range desired {
		if err := watcher.Add(path); err != nil {
			return fmt.Errorf("add watch %q: %w", path, err)
		}
	}
	return nil
}

func watchEvents(watcher *fsnotify.Watcher, rebuild func()) {
	timer := time.NewTimer(watchDebounce)
	timer.Stop()
	defer timer.Stop()

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if !event.Has(fsnotify.Write) && !event.Has(fsnotify.Create) && !event.Has(fsnotify.Remove) && !event.Has(fsnotify.Rename) {
				continue
			}
			Log(event.Name)
			timer.Reset(watchDebounce)
		case <-timer.C:
			rebuild()
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
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
	if err := configureWatcher(newWatcher); err != nil {
		if closeErr := newWatcher.Close(); closeErr != nil {
			Warn(closeErr.Error())
		}
		Fatal(err.Error())
	}
	watcher = newWatcher
	go watchEvents(newWatcher, func() {
		ParseGlobalConfigWrap(rootPath, true)
		if globalConfig == nil || themeConfig == nil {
			Warn("Parse config.yml failed; waiting for another change")
			return
		}
		if err := configureWatcher(newWatcher); err != nil {
			Warn(err.Error())
			return
		}
		Build()
		notifyReloadClients()
	})
}

func Websocket(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if c, err := upgrader.Upgrade(w, r, nil); err != nil {
		Warn(err)
	} else {
		reloadMu.Lock()
		reloadClients[c] = struct{}{}
		reloadMu.Unlock()
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
