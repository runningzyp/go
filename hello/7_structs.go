package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}
type Employee struct {
	Information Person
	ManagerID   int
}

type PersonJson struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type EmployeeJson struct {
	PersonJson
	ManagerID int
}

type ContractorJson struct {
	PersonJson
	CompanyID int
}

func main() {
	employee := Employee{ManagerID: 2, Information: Person{1, "Steve", "Jobs", "shanghai"}}
	fmt.Println(employee)
	fmt.Println(employee.Information.FirstName)

	employeeCopy := &employee
	employeeCopy.Information.FirstName = "Park"
	employeeCopy2 := employee
	employeeCopy2.Information.FirstName = "John"

	fmt.Println(employee)
	fmt.Println(*employeeCopy)

	employees := []EmployeeJson{
		EmployeeJson{
			PersonJson: PersonJson{
				LastName: "Doe", FirstName: "John",
			},
		},
		EmployeeJson{
			PersonJson: PersonJson{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}
	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []EmployeeJson
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
