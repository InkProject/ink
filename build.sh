#!/bin/sh

# binary build script for ink

build () {
  echo "building for $1 $2..."
  GOOS=$1 GOARCH=$2 go build -o release/ink
  cd release
  if [ $1 == "linux" ]
  then
    tar cvf - blog/* ink | gzip -9 - > ink_$1_$2.tar.gz
  else
    7za a -tzip -r ink_$1_$2.zip blog ink
  fi
  rm -rf ink
  cd ..
}

rm -rf release
mkdir -p release

rsync -av template/* release/blog --delete --exclude public --exclude theme/node_modules

build linux 386
build linux amd64
build darwin amd64
build windows 386
build windows amd64

rm -rf release/blog
