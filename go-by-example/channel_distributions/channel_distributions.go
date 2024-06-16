package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func main() {
	pings: make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}