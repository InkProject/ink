FROM golang:1.17
RUN go get github.com/InkProject/ink
EXPOSE 8000
WORKDIR /go/src/github.com/InkProject/ink
CMD ["ink", "preview", "template"]
