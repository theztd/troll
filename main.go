/*
  Hello world web application (example)
*/
package main

import (
	"flag"
)

func main() {
	// declare arguments
	flag.StringVar(&NAME, "name", "troll", "Define custom application name")
	flag.IntVar(&WAIT, "wait", 0, "Minimal wait time before each request")
	flag.StringVar(&DOC_ROOT, "root", "./public", "Define document root for serving files")
	flag.Parse()

	// it is better to be configurable via env
	ADDRESS = getEnv("ADDRESS", ":8080")

	// It is enought
	getRoutes()
	router.Run(ADDRESS)
}
