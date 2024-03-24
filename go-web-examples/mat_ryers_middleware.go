
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func createNewMiddleware() Middleware {
	// Create a new Middleware
	middleware := func(next http.HandleFunc) http.HandleFunc {

		// Define the http.HandlerFunc which is called by the server eventually
		handler := func(w http.ResponseWriter, r *http.Request) {
			
			// Do mddleware things

			// Call the next middleware/handler in chain
			next(w, r)
		}

		// Return newly created handler
		return handler
	}

	// Return newly created middleware
	return middleware
}

type Middleware func(http.HandleFunc) http.HandleFunc

// Logging logs all requests with its path and the time it took to process

func Logging() Middleware {

	// Create a new Middleware
	return func(f http.HandleFunc) http.HandleFunc{
		
		//  Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request){

			//  Do middleware things
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start))}()

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}

// Method ensures that url can only be requested with a specific method, else returns a 400 Bad Request
func Methods(m string) Middleware{

	// Create a new Middleware
	return func(f http.HandleFunc) http.HandleFunc {

		// Define the http.HandlerFunc
		return func(f http.ResponseWriter, r *http.Request) {

			// Do middleware things
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}


// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandleFunc, middleware ...Middleware) http.HandleFunc {
	for _, m := range middleware {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func main(){
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8080", nil)
}