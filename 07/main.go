package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
	// http.Handle("/favicon.ico", http.NotFoundHandler())
	// http.Handle("/public",http.StripPrefix("/public", http.FileServer(http.Dir("public"))))

}
