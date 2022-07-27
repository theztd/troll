/*
  Hello world web application (example)
*/
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// declare arguments
	flag.StringVar(&NAME, "name", "troll", "Define custom application name")
	flag.IntVar(&WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&DOC_ROOT, "root", "./public", "Define document root for serving files")
	flag.Parse()

	// it is better to be configurable via env
	PORT = getEnv("PORT", ":8080")

	log.Println("Starting web server", VERSION)
	log.Println(" - listen on port: ", PORT)
	log.Println("")

	// name := flag.String("name", "goapi")

	router := mux.NewRouter()

	router.HandleFunc("/{url:.*}.json", hello_json)
	router.HandleFunc("/{url:.*}", hello_txt)

	http.Handle("/", router)

	log.Println("Ready for connections...")

	log.Fatal(http.ListenAndServe(PORT, nil))

	log.Println("Server is down...")
}
