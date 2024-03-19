#!/bin/bash
nohup ./bin/$1 &
ps -ef | grep $1