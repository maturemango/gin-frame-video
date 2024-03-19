#!/bin/bash
echo "go build"
go build -o ./bin/$1
echo "build success"
chmod +x ./bin/$1
echo "chmod success"

#在项目初始路径下运行脚本文件(不管脚本文件在项目中的哪个路径下)