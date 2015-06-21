package main

type Lang map[string]struct {
	en string
	zh string
}

var lang Lang

func InitLang() {
	lang = Lang{
		"archive": {
			en: "Archive",
			zh: "归档",
		},
		"tag": {
			en: "Tag",
			zh: "标签",
		},
		"articles": {
			en: "Articles",
			zh: "篇文章",
		},
		"updated": {
			en: "updated",
			zh: "更新于",
		},
		"prev_page": {
			en: "Prev Page",
			zh: "上一页",
		},
		"next_page": {
			en: "Next Page",
			zh: "下一页",
		},
	}
}
