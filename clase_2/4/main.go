package main
import "fmt"

func main() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println("Edad Benjamin:", employees["Benjamin"] )
	fmt.Println("Cantidad de empleados mayores a 21:", EmployeesOlderThan(employees, 21))
	employees["Federico"] = 25
	delete(employees, "Pedro")
}

func EmployeesOlderThan(employees map[string]int, ageLimit int) (employeeCounter int){
	for _, age := range employees {
		if age > ageLimit {
			employeeCounter++
		}
	}
	return 
}