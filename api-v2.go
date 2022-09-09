package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type YamlStruct struct {
	Name        string
	Description string
	Version     string
	Endpoints   []Endpoint
}

type Endpoint struct {
	Method   string
	Path     string
	Code     int
	Response string
}

func loadYaml(path string) YamlStruct {
	yamlF, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	data := YamlStruct{}
	err2 := yaml.Unmarshal(yamlF, &data)
	if err2 != nil {
		log.Fatalln(err2)
	}

	return data
}

func v2RoutesAdd(rtG *gin.RouterGroup) {
	r := rtG.Group("/")
	log.Println("Loading V2 routes...")

	r.Use(requestid.New())

	r.GET("/info", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"version":    "0.0.1",
			"app_name":   "troll-dymanic-api",
			"client_ip":  c.ClientIP(),
			"referer":    c.Request.Referer(),
			"user-agent": c.Request.UserAgent(),
			"reqId":      requestid.Get(c),
		})
	})

	// if v2 yaml configuration exists, generate endpoints
	if _, err := os.Stat(V2_PATH); err == nil {
		cfg := loadYaml(V2_PATH)

		for _, x := range cfg.Endpoints {

			switch method := x.Method; method {
			case "GET":
				r.GET(x.Path, func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"version":    "0.0.1",
						"app_name":   "troll-dymanic-api",
						"client_ip":  c.ClientIP(),
						"referer":    c.Request.Referer(),
						"user-agent": c.Request.UserAgent(),
						"reqId":      requestid.Get(c),
						"msg":        x.Response,
					})
				})

			case "POST":
				r.POST(x.Path, func(c *gin.Context) {
					c.JSON(http.StatusOK, gin.H{
						"version":    "0.0.1",
						"app_name":   "troll-dymanic-api",
						"client_ip":  c.ClientIP(),
						"referer":    c.Request.Referer(),
						"user-agent": c.Request.UserAgent(),
						"reqId":      requestid.Get(c),
						"msg":        x.Response,
					})
				})

			default:
				log.Println("Skip, because method has not been defined " + x.Path)

			}
		}
	} else {
		log.Println("ERR: Unable to find file " + V2_PATH)
	}
}
