package main

import (
	"fmt"
)

func main() {
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)
	for val >= 0 {
		fmt.Println("You entered:", val)
		fmt.Scanf("%d", &val)
	}
	panic("输入负数")

}
