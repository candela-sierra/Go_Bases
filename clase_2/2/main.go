package main
import "fmt"

func main() {
	age := 23
	isEmployed := true
	seniority  := 10
	salary := 2000


	if age > 22 && isEmployed && seniority > 1 {
		if(salary > 100000) {
			fmt.Println("Podemos ofrecerle un prestamo sin interes")
		} else {
			fmt.Println("Podemos ofrecerle un prestamo con interes")
		}
	} else {
		fmt.Println("No podemos ofrecerle un prestamo")
	}
	
}