package main

import (
	"fmt"
	"math"
)

func main() {
	var firstName string
	var secondName, address string
	var age int
	const HTTPStatusOk = 200
	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
	var integer8 int8 = 127
	rune := 'G'
	//var unsign1 uint= -2
	//fmt.Println(unsign1)
	var boolFlag bool = true
	println(boolFlag)
	println(rune)
	println(math.MaxFloat32,math.MaxFloat64)
	fmt.Println(firstName, secondName, address, age)
	fmt.Println(StatusOK, StatusConnectionReset, StatusOtherError, HTTPStatusOk)
	fmt.Println(integer8)
}
