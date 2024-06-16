package main

import "fmt"

func main() {

	bool

	string

	int  int8  int16  int32  int64
	uint  uint8  uint16  uint32  uint64

	byte // alias for uint8

	rune // alias for int32. Represents a Unicode code point

	float32 float64

	complex64 complex128

	var smsSendingLimit int = 2
	var costPerSMS float32 = 1.85
	var hasPermission bool = false
	var username string = "Dr. Biko"


	fmt.Println("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}