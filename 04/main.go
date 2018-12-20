package main

import (
	"io"
	"net/http"
)

type dog int
type cat int

func (m dog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy")
}

func (m cat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "kitty")
}

func main() {
	var d dog
	var e cat

	// mux := http.NewServeMux()
	// mux.Handle("/dog/", d)
	// mux.Handle("/cat/", e)
	// http.ListenAndServe(":8080", mux)

	http.Handle("/dog/", d)
	http.Handle("/cat/", e)
	http.ListenAndServe(":8080", nil)
}
