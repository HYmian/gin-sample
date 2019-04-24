#! /bin/bash

log_dir="./logs"
if [ ! -d "$log_dir" ]; then
	mkdir logs
fi

#export MYSQL_HOST=192.168.0.101
#export MYSQL_PORT=3306
#export MYSQL_DATABASE=wise2c
#export MYSQL_USER=mian
#export MYSQL_PASSWD=admin

exec ./gin-sample -alsologtostderr=true -log_dir=./logs -v=2