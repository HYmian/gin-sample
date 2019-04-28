package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
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

		if err := orm.RegisterDataBase("default",
			"mysql",
			fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8",
				user,
				passwd,
				host,
				database,
			),
			30,
		); err != nil {
			glog.Fatalf("connect to mysql error: %s", err.Error())
		}

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

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select title, item from page").Values(&maps)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	pages := &Pages{}
	for i := int64(0); i < num; i++ {
		p := &Page{
			Title: maps[i]["title"].(string),
			Item:  maps[i]["item"].(string),
		}

		pages.Pages = append(pages.Pages, p)
	}

	c.HTML(http.StatusOK, "index", pages)
}

func PostPage(c *gin.Context) {
	pages := &Pages{}
	glog.Info(c.Request.RemoteAddr)

	o := orm.NewOrm()
	for _, i := range pages.Pages {
		if _, err := o.
			Raw("insert into page(title, item) values(?, ?)").
			SetArgs(i.Title, i.Item).
			Exec(); err != nil {
			glog.Errorf("insert page (%s, %s) error", i.Title, i.Item)
			c.String(http.StatusInternalServerError, "insert page error")
		}

	}

	c.Status(http.StatusOK)
}
