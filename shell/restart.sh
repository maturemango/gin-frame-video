#!/bin/bash
echo "kill current pid"
kill -INT $1
echo "restart gin-frame"
nohup ./bin/$2 & #需要忽略所有输出信息时改为 nohup ./bin/$2 >/dev/null 2>&1 &
ps -ef | grep $2