package main

import "fmt"

// Default Representation
s := fmt.Sprintf("I am %v years old", 10)

s := fmt.Sprintf("I am %v years old", "way too many")

// String
s := fmt.Sprintf("I am %s years old", "way too less")

// Integer
s := fmt.Sprintf("I am %d years old", 10)

// Float
s := fmt.Sprintf("I am %f years old", 10.523)
s := fmt.Sprintf("I am %.2f years old", 10.543)


func main(){
	const name = "Biko Steven"
	const openRate = 30.5
	var msg string = fmt.Sprintf("Hi %s, your open rate is %f percent%v", name, openRate \n)

	fmt.Print(msg)
}