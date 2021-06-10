package main

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

var (
	allowTrafficIncrease bool = false
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	defer glog.Flush()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/stress/:value", GetStress)
	r.GET("/ip", GetIP)
	r.POST("/flagger/traffic-increase", flaggerTrafficIncrease)
	r.PUT("/flagger/traffic-increase", AllowTrafficIncrease)

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
		bs[i] = byte(rand.Intn(94) + 33)
	}

	c.String(http.StatusOK, string(bs))
}

func GetIP(c *gin.Context) {
	c.String(http.StatusOK, c.Request.RemoteAddr)
}

func flaggerTrafficIncrease(c *gin.Context) {
	if allowTrafficIncrease {
		c.Status(http.StatusOK)
	} else {
		c.Status(http.StatusForbidden)
	}
}

func AllowTrafficIncrease(c *gin.Context) {
	allowTrafficIncrease = (c.Query("allow") == "true")
}
