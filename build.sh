#!/bin/bash

echo "Building..."

function BUILD() {
  echo "Building ${1} for linux arm64..."
  GOOS=linux GOARCH=arm64 go build -o build/${1}-linux-arm64

  echo "Building ${1} for linux amd64..."
  GOOS=linux GOARCH=amd64 go build -o build/${1}-linux-amd64

  echo "Building ${1} for darwin arm64..."
  GOOS=darwin GOARCH=arm64 go build -o build/${1}-darwin-arm64
}

cd server
BUILD "server"
cd - > /dev/null

cd sync
BUILD "sync"
cd - > /dev/null

echo "Finished!"