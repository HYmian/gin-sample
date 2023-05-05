module github.com/HYmian/gin-sample

go 1.16

require (
	github.com/gin-gonic/gin v1.9.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190422183909-d864b10871cd
	golang.org/x/net => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190424160641-4347357a82bc
	golang.org/x/text => github.com/golang/text v0.3.2-0.20190424151008-b1379a7b4714
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190424031103-cb2dda6eabdf
	google.golang.org/appengine => github.com/golang/appengine v1.5.0
)
