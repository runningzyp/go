package main

import (
	"example.com/greetings"
	"fmt"
	"math"
)

func main() {
	// Get a greeting message and print it.
	message,_ := greetings.Hello("Gladys")
	fmt.Println(message)
	ret := math.Abs(-1.23)
	fmt.Println(ret)
}
