package main

import (
	"strconv"
)

sum := sum(os.Args[1], os.Args[2])
	println("Sum:", sum)
	print(rand.Int())
}
func sum(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}
