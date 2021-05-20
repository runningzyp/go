package main

import "fmt"

func main()  {
	var firstName string
	var secondName,address string
	var age int
	const HTTPStatusOk=200
	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
	fmt.Println(firstName,secondName,address,age)
	fmt.Println(StatusOK,StatusConnectionReset,StatusOtherError,HTTPStatusOk)

}