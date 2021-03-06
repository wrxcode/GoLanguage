package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

var employees [5]Employee

func EmployeeByID(id int) Employee {

	for _, e := range employees {
		if e.ID == id {
			return e
		}
	}
	return dilbert
}

func main() {

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
	id := dilbert.ID

	e := EmployeeByID(id)
	e.Salary = 0

	EmployeeByID(id).Salary = 0
}
