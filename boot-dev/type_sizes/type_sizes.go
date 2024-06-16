package main

import "fmt"

func main(){
	int int8 int16 int32 int64 // whole numbers

	uint uint8 uint16 uint32 uint64 // positive whole numbers

	float32 float64 // decimal numbers

	complex64 complex128 // imaginary numbers (rare)

	temperatureFloat := 88.26
	temperatureInt := int64(temperatureFloat)

	accountAge := 2.6
	accountAgeInt := int64(accountAge)

	fmt.Println("Your account has existed for", accountAgeInt, "years")
}