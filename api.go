package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/InkProject/ink.go"
	"github.com/facebookgo/symwalk"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	// "fmt"
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

func replyJSON(ctx *ink.Context, status int, data interface{}) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		http.Error(ctx.Res, err.Error(), http.StatusInternalServerError)
		ctx.Stop()
		return
	}
	if status == http.StatusOK {
		ctx.Header().Set("Content-Type", "application/json")
		ctx.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Res.Write(jsonStr)
	} else {
		Warn(data)
		http.Error(ctx.Res, data.(string), status)
	}
	ctx.Stop()
}

func UpdateArticleCache() {
	articleCache = make(map[string]CacheArticleInfo, 0)
	symwalk.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		fileExt := strings.ToLower(filepath.Ext(path))
		if fileExt == ".md" {
			fileName := strings.TrimPrefix(strings.TrimSuffix(strings.ToLower(path), ".md"), "template/source/")
			config, _ := ParseArticleConfig(path)
			id := hashPath(path)
			articleCache[string(id)] = CacheArticleInfo{
				Name:    fileName,
				Path:    path,
				Date:    ParseDate(config.Date),
				Article: config,
			}
		}
		return nil
	})
}

func ApiListArticle(ctx *ink.Context) {
	UpdateArticleCache()
	replyJSON(ctx, http.StatusOK, articleCache)
}

func ApiGetArticle(ctx *ink.Context) {
	UpdateArticleCache()
	article, ok := articleCache[ctx.Param["id"]]
	if !ok {
		replyJSON(ctx, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.Path
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, string(data))
}

func ApiRemoveArticle(ctx *ink.Context) {
	UpdateArticleCache()
	article, ok := articleCache[ctx.Param["id"]]
	if !ok {
		replyJSON(ctx, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.Path
	err := os.Remove(filePath)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, nil)
}

func ApiCreateArticle(ctx *ink.Context) {
	decoder := json.NewDecoder(ctx.Req.Body)
	var article NewArticle
	err := decoder.Decode(&article)
	if err != nil {
		replyJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	filePath := filepath.Join(sourcePath, article.Name+".md")
	err = ioutil.WriteFile(filePath, []byte(article.Content), 0644)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, map[string]string{
		"id": hashPath(filePath),
	})
}

func ApiSaveArticle(ctx *ink.Context) {
	UpdateArticleCache()
	decoder := json.NewDecoder(ctx.Req.Body)
	var article OldArticle
	err := decoder.Decode(&article)
	if err != nil {
		replyJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	cacheArticle, ok := articleCache[ctx.Param["id"]]
	if !ok {
		replyJSON(ctx, http.StatusNotFound, "Not Found")
		return
	}
	// Write
	path := cacheArticle.Path
	err = ioutil.WriteFile(path, []byte(article.Content), 0644)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, nil)
}

func getFormFile(ctx *ink.Context, field string) (data []byte, handler *multipart.FileHeader, err error) {
	file, handler, err := ctx.Req.FormFile(field)
	if err != nil {
		replyJSON(ctx, http.StatusBadRequest, err.Error())
		return nil, handler, err
	}
	data, err = ioutil.ReadAll(file)
	if err != nil {
		replyJSON(ctx, http.StatusBadRequest, err.Error())
		return data, handler, err
	}
	return data, handler, err
}

func ApiUploadFile(ctx *ink.Context) {
	UpdateArticleCache()
	fileData, handler, err := getFormFile(ctx, "file")
	if err != nil {
		replyJSON(ctx, http.StatusBadRequest, err.Error())
		return
	}
	articleId := ctx.Req.FormValue("article_id")
	article, ok := articleCache[articleId]
	if !ok {
		replyJSON(ctx, http.StatusNotFound, "Not Found")
		return
	}
	fileDirPath := filepath.Join(sourcePath, "images", article.Name)
	err = os.MkdirAll(fileDirPath, 0777)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if err = ioutil.WriteFile(filepath.Join(fileDirPath, handler.Filename), fileData, 0777); err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, map[string]string{
		"path": "-/" + filepath.Join("images", article.Name, handler.Filename),
	})
}

func ApiGetConfig(ctx *ink.Context) {
	filePath := filepath.Join(rootPath, "config.yml")
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, string(data))
}

func ApiSaveConfig(ctx *ink.Context) {
	content, err := ioutil.ReadAll(ctx.Req.Body)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	filePath := filepath.Join(rootPath, "config.yml")
	err = ioutil.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, nil)
}

// func ApiRenameArticle(ctx *ink.Context) {
// 	// Rename
// 	cacheArticle, ok := articleCache[ctx.Param["id"]]
// 	if !ok {
// 		replyJSON(ctx, http.StatusNotFound, "Not Found")
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
