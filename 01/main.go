package main

import (
	"fmt"
	"net/http"
)

type abc int

func (m abc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "abc response ...")
}

func main() {
	var d abc
	http.ListenAndServe(":8080", d)
}
