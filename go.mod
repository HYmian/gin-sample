module github.com/HYmian/gin-sample

require (
	github.com/astaxie/beego v1.11.1
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/martini-contrib/binding v0.0.0-20160701174519-05d3e151b6cf
	github.com/martini-contrib/render v0.0.0-20150707142108-ec18f8345a11
	github.com/oxtoacart/bpool v0.0.0-20190227141107-8c4636f812cc // indirect
	google.golang.org/appengine v1.5.0 // indirect
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
