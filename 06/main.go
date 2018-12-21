package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/abc", abc)
	http.ListenAndServe(":8080", nil)
}

func abc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// io.WriteString(w,``)
}
