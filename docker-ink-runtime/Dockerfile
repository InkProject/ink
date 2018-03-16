# 本 Dockerfile 用来构建 ink-runtime 为基础镜像,以便使用.

# 基于
# 生成 ink 的可执行文件
FROM golang:alpine AS ink-build
RUN apk add --update git
RUN go get -u -v github.com/taadis/ink
WORKDIR /go/src/github.com/taadis/ink/
RUN go install -v
WORKDIR /go/bin/
RUN ls -l

# 基于
# 生成最小化的 ink 基础镜像.
# 使用过程中只需要基于镜像 taadis/ink-runtime 即可.
FROM alpine:latest AS ink-runtime
WORKDIR /ink/
COPY --from=ink-build /go/bin/ink ./
COPY --from=ink-build /go/src/github.com/taadis/ink/template/ ./template/
RUN ls -l
EXPOSE 8000
CMD ["./ink","serve","template"]
