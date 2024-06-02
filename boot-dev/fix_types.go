package main

import "fmt"

func main(){
	var senderName string = "Syl"
	var recipient string = "Kaladin"
	var message string = "The words, Kaladin. You have to speak the Words!"

	fmt.Println("%s to %s: %\n", senderName, recipient, message)
}