package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

type abc int

func (m abc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "abc response ...")
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func main() {
	var d abc
	http.ListenAndServe(":8080", d)
}
