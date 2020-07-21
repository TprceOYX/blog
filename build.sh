#!/bin/bash

RUN_NAME="blog"
mkdir -p output
export GO111MODULE=on
go mod download

go build -o output/${RUN_NAME}
chmod +x output/${RUN_NAME}