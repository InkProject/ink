FROM golang:1.4.2
go get github.com/InkProject/ink
ink preview $GOPATH/src/github.com/InkProject/ink/template
