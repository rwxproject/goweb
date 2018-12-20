package main

import (
	"io"
	"net/http"
)

type abc int

func (m abc) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		io.WriteString(res, "doggy")
	case "/cat":
		io.WriteString(res, "kitty")
	}
}

func main() {
	var d abc
	http.ListenAndServe(":8080", d)
}
