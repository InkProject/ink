#!/bin/sh

# build binary releases for ink

build () {
  echo "building for $1 $2..."
  suffix=""
  if [ $1 = "windows" ]
  then
    suffix=".exe"
  fi
  GOOS=$1 GOARCH=$2 go build -o release/ink$suffix
  cd release
  if [ $1 = "linux" ]
  then
    tar cvf - blog/* ink$suffix | gzip -9 - > ink_$1_$2.tar.gz
  else
    7z a -tzip -r ink_$1_$2.zip blog ink$suffix
  fi
  rm -rf ink$suffix
  cd ..
}

rm -rf release
mkdir -p release

rsync -av template/* release/blog --delete --exclude public --exclude theme/node_modules

build linux 386
build linux amd64
build linux arm
build linux arm64

build darwin amd64
build darwin arm64

build windows 386
build windows amd64
build windows arm
build windows arm64

rm -rf release/blog
