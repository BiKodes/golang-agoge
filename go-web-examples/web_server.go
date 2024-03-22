package main

import (
	"fmt"
	"net/http"
)


// Registering a Request Handler
http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Jambo, you have requested: %s\n", r.URL.Path)
})

// Listen for HTTP Connections
http.ListenAndServer("80", nil)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
    })

    http.ListenAndServe(":80", nil)
}