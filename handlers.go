package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ret struct {
	status  int
	name    string
	version string
	url     string
	delay   int
}

type Map map[string]interface{}

func hello_txt(w http.ResponseWriter, r *http.Request) {
	log.Println(" Requested url ", r.URL.Path)

	// parse params from request
	params := mux.Vars(r)

	// Generate delay
	delay := time.Millisecond * time.Duration(WAIT+rand.Intn(500))
	time.Sleep(delay)

	// generate web response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "url: /%s\nname: %s\nversion: %s\nmin_delay: %d", params["url"], NAME, VERSION, delay)
}

func hello_json(w http.ResponseWriter, r *http.Request) {
	log.Println(" Requested url ", r.URL.Path)

	params := mux.Vars(r)

	// Generate delay
	delay := time.Millisecond * time.Duration(WAIT+rand.Intn(500))
	time.Sleep(delay)

	data := Map{
		"name":    NAME,
		"url":     params["url"],
		"version": VERSION,
		"delay":   int(delay),
		"status":  200,
	}
	log.Println(data)

	// data, _ := json.Marshal(ret{
	// 	status:  200,
	// 	name:    NAME,
	// 	url:     params["url"],
	// 	version: VERSION,
	// 	delay:   int(delay),
	// })
	// fmt.Println(data)

	// generate web response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	// fmt.Fprintf(w, "url: /%s\nname: %s\nversion: %s\nmin_delay: %d", r.URL.Path[1:], NAME, VERSION, delay)
}
