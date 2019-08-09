#! /bin/sh

x=`curl http://gin-sample.test.svc.cluster.local:8080/stress/3 -w '%{size_download}' -so /dev/null`;
if [ $x -eq '3072' ]; then
    exit 0;
else
    echo x=$x;
    exit 1;
fi
