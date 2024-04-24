#!/bin/bash

if [ ! -d "./bin" ]; then
  echo "bin dir does not exist, making it now"
  mkdir bin
else
  echo "bin dir already exists" 
fi

build_server() {
  go build -o ./bin/server cmd/server/main.go
}

build_client() {
  go build -o ./bin/client cmd/client/main.go
}

if [ "$1" == "server" ]; then build_server
elif [ "$1" == "client" ]; then build_client 
elif [ "$1" == "all" ]; then build_server; build_client
fi
