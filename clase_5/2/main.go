package main

import "fmt"

type Person struct {
	ID          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	ID       int
	Position string
	Person   Person
}

func main() {
	newPerson := Person{Name: "Jorge", ID: 2, DateOfBirth: "07/02/1960"}
	newEmployee := Employee{Person: newPerson, ID: 45, Position: "Cajero"}
	newEmployee.printEmployee()
}

func (Employee Employee) printEmployee() {
	fmt.Println("Employee ID:", Employee.ID)
	fmt.Println("Position:", Employee.Position)
	fmt.Println("Persona ID:", Employee.Person.ID)
	fmt.Println("Date of Birth:", Employee.Person.DateOfBirth)
	fmt.Println("Name:", Employee.Person.Name)
}
