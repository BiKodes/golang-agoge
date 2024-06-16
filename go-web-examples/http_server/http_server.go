package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Karibu to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	htpp.Handle("/static/", htpp.StripPrefix("/static/", fs))

	htpp.ListenAndServe(":80", nil)
}
