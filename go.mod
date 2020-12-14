module github.com/InkProject/ink

go 1.15

replace (
	github.com/InkProject/blackfriday v0.0.0-20181012080017-b70c36859218 => github.com/russross/blackfriday v1.6.0
)

require (
	github.com/InkProject/blackfriday v0.0.0-20181012080017-b70c36859218
	github.com/InkProject/ink.go v0.0.0-20160120061933-86de6d066e8d
	github.com/urfave/cli/v2 v2.1.1
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/facebookgo/ensure v0.0.0-20200202191622-63f1cf65ac4c // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20200203212716-c811ad88dec4 // indirect
	github.com/facebookgo/symwalk v0.0.0-20150726040526-42004b9f3222
	github.com/facebookgo/testname v0.0.0-20150612200628-5443337c3a12 // indirect
	github.com/fsnotify/fsnotify v1.4.9
	github.com/gorilla/feeds v1.1.1
	github.com/gorilla/websocket v1.4.2
	github.com/kr/pretty v0.2.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)
