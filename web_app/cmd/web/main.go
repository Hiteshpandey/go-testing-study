package main

import (
	"log"
	"net/http"
)

type application struct {
}

func main() {
	// set up configuration
	var app application
	// setup routes
	mux := app.routes()
	// print message
	log.Println("Starting server at port 8080")
	// start server
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
