package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/InkProject/ink.go"
	"github.com/facebookgo/symwalk"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type NewArticle struct {
	Name    string
	Content string
}

type OldArticle struct {
	Content string
}

var articleCache map[string]interface{}

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
	articleCache = make(map[string]interface{}, 0)
	symwalk.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		fileExt := strings.ToLower(filepath.Ext(path))
		if fileExt == ".md" {
			fileName := strings.TrimSuffix(strings.ToLower(filepath.Base(path)), ".md")
			config, _ := ParseArticleConfig(path)
			md5Hex := md5.Sum([]byte(path))
			id := hex.EncodeToString(md5Hex[:])
			articleCache[string(id)] = map[string]interface{}{
				"name":    fileName,
				"path":    path,
				"article": config,
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
	filePath := article.(map[string]interface{})["path"].(string)
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
	filePath := article.(map[string]interface{})["path"].(string)
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
	replyJSON(ctx, http.StatusOK, nil)
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
	path := cacheArticle.(map[string]interface{})["path"].(string)
	err = ioutil.WriteFile(path, []byte(article.Content), 0644)
	if err != nil {
		replyJSON(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	replyJSON(ctx, http.StatusOK, nil)
}


// Rename
// cacheArticle, ok := articleCache[ctx.Param["id"]]
// if !ok {
// 	replyJSON(ctx, http.StatusNotFound, "Not Found")
// 	return
// }
// oldPath := cacheArticle.(map[string]interface{})["path"].(string)
// newPath := filepath.Join(sourcePath, newArticle.Name+".md")
// err = os.Rename(oldPath, newPath)
// if err != nil {
// 	replyJSON(ctx, http.StatusInternalServerError, err.Error())
// 	return
// }
