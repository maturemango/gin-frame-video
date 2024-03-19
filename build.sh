#!/bin/bash
echo "go build"
go build -o ./bin/$1
echo "build success"
chmod +x ./bin/$1
echo "chmod success"