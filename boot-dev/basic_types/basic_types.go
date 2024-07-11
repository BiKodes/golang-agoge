package main

import "fmt"

func main() {

	// byte // alias for uint8

	// rune // alias for int32. Represents a Unicode code point

	var smsSendingLimit int = 2
	var costPerSMS float32 = 1.85
	var hasPermission bool = false
	var username string = "Dr. Biko"

	fmt.Println("%v, %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
