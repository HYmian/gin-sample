package main

import (
	"flag"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	_ "github.com/go-sql-driver/mysql"
)

type Pages struct {
	Pages []*Page `json:"pages" binding:"required"`
}

type Page struct {
	Title string `json:"title" binding:"required"`
	Item  string `json:"item" binding:"required"`
}

var (
	havedb = flag.Bool("havedb", false, "demo connect to a DB if true")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	defer glog.Flush()
	flag.Parse()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if *havedb {
		host := "galera-lb"
		user := os.Getenv("MYSQL_USER")
		passwd := os.Getenv("MYSQL_PASSWORD")
		database := os.Getenv("MYSQL_DATABASE")

		r.GET("/", GetPages)
		r.POST("/", PostPage)
	}
	r.GET("/stress/:value", GetStress)

	r.Run("0.0.0.0:8080")
}

func GetStress(c *gin.Context) {
	v := c.Param("value")
	iv, err := strconv.ParseUint(v, 10, 64)

	if err != nil {
		iv = 50 * 1024
	} else {
		iv = iv * 1024
	}

	bs := make([]byte, iv)
	for i := uint64(0); i < iv; i++ {
		bs[i] = byte(rand.Intn(95) + 32)
	}

	c.String(http.StatusOK, string(bs))
}

func GetPages(c *gin.Context) {
	glog.Info(c.Request.RequestURI)

	c.Status(http.StatusOK)
}

func PostPage(c *gin.Context) {
	pages := &Pages{}
	glog.Info(c.Request.RemoteAddr)

	c.Status(http.StatusOK)
}
