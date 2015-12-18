package main

import (
	"encoding/json"
	"github.com/ant0ine/go-json-rest/rest"
	"io/ioutil"
	"os"
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"path/filepath"
	"strings"
)

type NewArticle struct {
	Name    string
	Content string
}

var articleCache map[string]interface{}

func responseJSON(w rest.ResponseWriter, status int, data interface{}) {
	if status == http.StatusOK {
		w.WriteJson(data)
	} else {
		rest.Error(w, data.(string), status)
	}
}

func UpdateArticleCache() {
	articleCache = make(map[string]interface{}, 0)
	filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
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

func ApiListArticle(w rest.ResponseWriter, req *rest.Request) {
	UpdateArticleCache()
	responseJSON(w, http.StatusOK, articleCache)
}

func ApiGetArticle(w rest.ResponseWriter, req *rest.Request) {
	UpdateArticleCache()
	article, ok := articleCache[req.PathParam("id")]
	if !ok {
		responseJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.(map[string]interface{})["path"].(string)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, string(data))
}

func ApiRemoveArticle(w rest.ResponseWriter, req *rest.Request) {
	UpdateArticleCache()
	article, ok := articleCache[req.PathParam("id")]
	if !ok {
		responseJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	filePath := article.(map[string]interface{})["path"].(string)
	err := os.Remove(filePath)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, nil)
}

func ApiCreateArticle(w rest.ResponseWriter, req *rest.Request) {
	decoder := json.NewDecoder(req.Request.Body)
	var article NewArticle
	err := decoder.Decode(&article)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	filePath := filepath.Join(sourcePath, article.Name + ".md")
	err = ioutil.WriteFile(filePath, []byte(article.Content), 0644)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, nil)
}

func ApiModifyArticle(w rest.ResponseWriter, req *rest.Request) {
	UpdateArticleCache()
	decoder := json.NewDecoder(req.Request.Body)
	var newArticle NewArticle
	err := decoder.Decode(&newArticle)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, err.Error())
		return
	}
	// Rename
	cacheArticle, ok := articleCache[req.PathParam("id")]
	if !ok {
		responseJSON(w, http.StatusNotFound, "Not Found")
		return
	}
	oldPath := cacheArticle.(map[string]interface{})["path"].(string)
	newPath := filepath.Join(sourcePath, newArticle.Name + ".md")
	err = os.Rename(oldPath, newPath)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	// Write
	err = ioutil.WriteFile(newPath, []byte(newArticle.Content), 0644)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseJSON(w, http.StatusOK, nil)
}
