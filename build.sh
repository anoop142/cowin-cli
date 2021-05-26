#!/bin/bash

type setopt >/dev/null 2>&1

SOURCE_FILE=`echo $@ | sed 's/\.go//'`
CURRENT_DIRECTORY=${PWD##*/}
BASE_BIN_DIRECTORY="bin"
BASE_RELEASE_DIRECTORY="release"
SCRIPTS_DIRECTORY="scripts"
OUTPUT="cowin-cli"
PLATFORMS="darwin/amd64" 
PLATFORMS="$PLATFORMS darwin/arm64"
PLATFORMS="$PLATFORMS windows/amd64"
PLATFORMS="$PLATFORMS windows/386" 
PLATFORMS="$PLATFORMS linux/amd64"
BUILD_FLAGS="'-s -w'"

build(){

  for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    BIN_DIRECTORY="${BASE_BIN_DIRECTORY}/${GOOS}_${GOARCH}"
    mkdir -p "$BIN_DIRECTORY"
    BIN_FILENAME="${OUTPUT}"
    if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
    CMD="env CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_DIRECTORY/BIN_FILENAME} -ldflags ${BUILD_FLAGS}  $@"
    echo "${CMD}"
    eval $CMD
  done
  }

clean(){
  rm -rf bin release
  go clean
}

release(){
  mkdir -p "$BASE_RELEASE_DIRECTORY"
  for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    BIN_DIRECTORY="${BASE_BIN_DIRECTORY}/${GOOS}_${GOARCH}"
    BIN_FILENAME="${OUTPUT}"
    ZIP_NAME="${OUTPUT}_${GOOS}_${GOARCH}.zip"
    if [[ "${GOOS}" == "windows" ]]; then BIN_FILENAME="${BIN_FILENAME}.exe"; fi
    CMD="zip -j9 ${BASE_RELEASE_DIRECTORY}/${ZIP_NAME} ${BIN_DIRECTORY}/${BIN_FILENAME} ${SCRIPTS_DIRECTORY}/${GOOS}/*"
    echo $CMD
    eval $CMD

  done
}


case "$1" in 

    "build")
      build
      ;;
    "release")
      clean 
      build && release
      ;;
    "clean")
      clean
      ;;
      *)
      build
      ;;
  esac
