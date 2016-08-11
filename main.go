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
	"github.com/go-martini/martini"
	"github.com/golang/glog"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"

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
	flag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer())

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

		m.Get("/", GetPages)
		m.Post("/", binding.Bind(Pages{}), PostPage)
	}
	m.Get("/stress/:value", GetStress)

	m.Run()
}

func GetStress(param martini.Params) (int, string) {
	v := param["value"]
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

	return 200, string(bs)
}

func GetPages(req *http.Request, r render.Render) {
	glog.Info(req.RequestURI)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select title, item from page").Values(&maps)
	if err != nil {
		r.Text(500, err.Error())
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

	r.HTML(200, "index", pages)

}

func PostPage(req *http.Request, pages Pages) (int, string) {
	glog.Info(req.RemoteAddr)

	o := orm.NewOrm()
	for _, i := range pages.Pages {
		if _, err := o.
			Raw("insert into page(title, item) values(?, ?)").
			SetArgs(i.Title, i.Item).
			Exec(); err != nil {
			glog.Errorf("insert page (%s, %s) error", i.Title, i.Item)
			return 500, "insert page error"
		}

	}
	return 200, ""
}
