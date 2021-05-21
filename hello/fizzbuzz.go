package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%d :Fizzbuzz\n", i)
		case i%3 == 0:
			fmt.Printf("%d :Fizz\n", i)
		case i%5 == 0:
			fmt.Printf("%d :buzz\n", i)

		}
	}
}
