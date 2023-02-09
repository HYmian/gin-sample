module github.com/HYmian/gin-sample

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-playground/validator/v10 v10.6.1 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/sys v0.0.0-20210608053332-aa57babbf139 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
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
