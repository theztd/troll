package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Endpoint struct {
	Path     string
	Method   string
	Code     int
	Response string
}

func DynamicRouter(configFile string) {
	r := mux.NewRouter()
	log.Panicln("Loading config file " + configFile)

	conf := []Endpoint{
		{"/status", "GET", 200, "Application is up and running"},
		{"/info", "GET", 200, "Latest version of the application"},
	}

	for _, endpoint := range conf {
		log.Println("INIT: prepare route" + endpoint.Path)
		route := endpoint
		r.HandleFunc(route.Path, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			log.Println("INFO: Requested url " + route.Path + "from " + r.RemoteAddr)
			w.WriteHeader(route.Code)
			io.WriteString(w, route.Response)
		}).Methods(route.Method)
	}
}
