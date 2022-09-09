package main

import (
	"io/ioutil"
	"log"
	"net/http"

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

	cfg := loadYaml("./api_example.yaml")

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

	/*
		cfg := []Endpoint{
			{"/users", 200, "list of all users"},
			{"/machines", 200, "list of the all machines in our factory"},
		}
	*/
	for _, x := range cfg.Endpoints {
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
	}
}
