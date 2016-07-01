package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/go-martini/martini"
	"github.com/golang/glog"
	"github.com/martini-contrib/render"

	_ "github.com/go-sql-driver/mysql"
)

type Page struct {
	Title string
	Items []string
}

func main() {
	flag.Parse()

	m := martini.Classic()
	m.Use(render.Renderer())

	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "galera-lb"
	}
	port := "3306"
	database := "wise2c"
	user := "wise2c"
	passwd := "test"

	if err := orm.RegisterDataBase("default",
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			user,
			passwd,
			host,
			port,
			database,
		),
		30,
	); err != nil {
		glog.Fatalf("connect to mysql error: %s", err.Error())
	}

	m.Get("/", GetPages)

	m.Run()
}

func GetPages(req *http.Request, r render.Render) {
	glog.Info(req.RequestURI)

	var maps []orm.Params
	o := orm.NewOrm()
	num, err := o.Raw("select title, item from page limit 1").Values(&maps)
	if num > 0 {
		p := &Page{
			Title: maps[0]["title"].(string),
			Items: []string{
				maps[0]["item"].(string),
			},
		}

		r.HTML(200, "index", p)
	} else {
		glog.Error(err.Error())
	}
}
