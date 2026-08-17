package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type NewArticle struct {
	Name    string
	Content string
}

type OldArticle struct {
	Content string
}

type CacheArticleInfo struct {
	Name    string
	Path    string
	Date    time.Time
	Article *ArticleConfig
}

var articleCache map[string]CacheArticleInfo

func hashPath(path string) string {
	md5Hex := md5.Sum([]byte(path))
	return hex.EncodeToString(md5Hex[:])
}

func replyJSON(w http.ResponseWriter, status int, data any) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if status == http.StatusOK {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		if _, err := w.Write(jsonStr); err != nil {
			Warn(err.Error())
		}
	} else {
		Warn(data)
		http.Error(w, data.(string), status)
	}
}

func UpdateArticleCache() {
	articleCache = make(map[string]CacheArticleInfo, 0)
	if err := walkSymlinks(sourcePath, func(path string, _ os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fileExt := strings.ToLower(filepath.Ext(path))
		if fileExt == ".md" {
			fileName := strings.TrimPrefix(strings.TrimSuffix(strings.ToLower(path), ".md"), "template/source/")
			config, _ := ParseArticleConfig(path)
			if config == nil {
				return nil
			}
			id := hashPath(path)
			articleCache[id] = CacheArticleInfo{
				Name:    fileName,
				Path:    path,
				Date:    ParseDate(config.Date),
				Article: config,
			}
		}
		return nil
	}); err != nil {
		Warn(err.Error())
	}
}

func ApiListArticle(w http.ResponseWriter, r *http.Request) {
	UpdateArticleCache()
	replyJSON(w, http.StatusOK, articleCache)
}

func ApiGetArticle(w http.ResponseWriter, r *http.Request) {
	UpdateArticleCache()
	article, ok := articleCache[r.PathValue("id")]
	if !ok {
		replyJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.Path
	data, err := os.ReadFile(filePath)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, string(data))
}

func ApiRemoveArticle(w http.ResponseWriter, r *http.Request) {
	UpdateArticleCache()
	article, ok := articleCache[r.PathValue("id")]
	if !ok {
		replyJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.Path
	err := os.Remove(filePath)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, nil)
}

func ApiCreateArticle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var article NewArticle
	err := decoder.Decode(&article)
	if err != nil {
		replyJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	filePath := filepath.Join(sourcePath, article.Name+".md")
	err = os.WriteFile(filePath, []byte(article.Content), 0644)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, map[string]string{
		"id": hashPath(filePath),
	})
}

func ApiSaveArticle(w http.ResponseWriter, r *http.Request) {
	UpdateArticleCache()
	decoder := json.NewDecoder(r.Body)
	var article OldArticle
	err := decoder.Decode(&article)
	if err != nil {
		replyJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	cacheArticle, ok := articleCache[r.PathValue("id")]
	if !ok {
		replyJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	// Write
	path := cacheArticle.Path
	err = os.WriteFile(path, []byte(article.Content), 0644)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, nil)
}

func getFormFile(w http.ResponseWriter, r *http.Request, field string) (data []byte, handler *multipart.FileHeader, err error) {
	file, handler, err := r.FormFile(field)
	if err != nil {
		replyJSON(w, http.StatusBadRequest, err.Error())
		return nil, handler, err
	}
	data, err = io.ReadAll(file)
	if err != nil {
		replyJSON(w, http.StatusBadRequest, err.Error())
		return data, handler, err
	}
	return data, handler, err
}

func ApiUploadFile(w http.ResponseWriter, r *http.Request) {
	UpdateArticleCache()
	fileData, handler, err := getFormFile(w, r, "file")
	if err != nil {
		replyJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	articleID := r.FormValue("article_id")
	article, ok := articleCache[articleID]
	if !ok {
		replyJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	fileDirPath := filepath.Join(sourcePath, "images", article.Name)
	err = os.MkdirAll(fileDirPath, 0777)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err = os.WriteFile(filepath.Join(fileDirPath, handler.Filename), fileData, 0777); err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, map[string]string{
		"path": "-/" + filepath.Join("images", article.Name, handler.Filename),
	})
}

func ApiGetConfig(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join(rootPath, "config.yml")
	data, err := os.ReadFile(filePath)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, string(data))
}

func ApiSaveConfig(w http.ResponseWriter, r *http.Request) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	filePath := filepath.Join(rootPath, "config.yml")
	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		replyJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(w, http.StatusOK, nil)
}

// func ApiRenameArticle(w http.ResponseWriter, r *http.Request) {
// 	// Rename
// 	cacheArticle, ok := articleCache[r.PathValue("id")]
// 	if !ok {
// 		replyJSON(w, http.StatusNotFound, "Not Found")
// 		return
// 	}
// 	oldPath := cacheArticle.(map[string]CacheArticleInfo)["path"].(string)
// 	newPath := filepath.Join(sourcePath, newArticle.Name+".md")
// 	err = os.Rename(oldPath, newPath)
// 	if err != nil {
// 		replyJSON(ctx, http.StatusInternalServerError, err.Error())
// 		return
// 	}
// }
