package main

import (
	"html/template"
	"os"
	"path/filepath"
)

type FuncContext struct {
	rootPath   string
	themePath  string
	publicPath string
	currentCwd string
	global     *GlobalConfig
}

func (ctx FuncContext) FuncMap() template.FuncMap {
	return template.FuncMap{
		"i18n":     ctx.I18n,
		"readFile": ctx.ReadFile,
	}
}

func (ctx FuncContext) I18n(val string) string {
	return ctx.global.I18n[val]
}

func (ctx FuncContext) ReadFile(path string) template.HTML {
	bytes, _ := os.ReadFile(filepath.Join(ctx.currentCwd, path))
	return template.HTML(bytes)
}
