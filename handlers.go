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

type Data map[string]interface{}

func hello_txt(w http.ResponseWriter, r *http.Request) {
	req_id := int(time.Now().Unix()) * rand.Intn(99)
	log.Println(req_id, " Requested url ", r.URL.Path, "user-agent:", r.UserAgent())

	// parse params from request
	params := mux.Vars(r)
	fmt.Println(params)

	// Generate delay
	delay := time.Millisecond * time.Duration(WAIT+rand.Intn(500))
	time.Sleep(delay)

	// generate web response
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "url: /%s\nname: %s\nversion: %s\nrequest_id: %d\nmin_delay: %d", params["url"], NAME, VERSION, req_id, delay)
}

func hello_json(w http.ResponseWriter, r *http.Request) {
	req_id := int(time.Now().Unix()) + rand.Intn(999999)
	log.Println(req_id, "Requested url", r.URL.Path, "user-agent:", r.UserAgent())

	params := mux.Vars(r)

	/*
		parse random data from json request
	*/

	// unformated map will be in body_json variable
	var body_json Data
	raw_body := json.NewDecoder(r.Body)
	raw_body.Decode(&body_json)

	// formated json
	pretty, _ := json.MarshalIndent(body_json, "  ", "  ")

	// Print received data to log
	log.Println(req_id, "DATA:", string(pretty))
	log.Println(req_id, "End DATA")

	// Generate delay
	delay := time.Millisecond * time.Duration(WAIT+rand.Intn(500))
	time.Sleep(delay)

	data := Map{
		"name":       NAME,
		"url":        params["url"],
		"version":    VERSION,
		"delay":      int(delay),
		"request_id": req_id,
		"status":     200,
		"body":       body_json,
	}

	// generate web response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
	log.Println(req_id, "END request")
}
