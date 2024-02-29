package main
import "fmt"

func main() {
	var monthNumber  int = 7
	
	fmt.Println(GetMonthName(monthNumber))
	fmt.Println(GetMonthName2(monthNumber))

}

func GetMonthName(monthNumber int) (string){
	if (monthNumber < 1 || monthNumber > 12){
		return "Mes invalido"
	}

	months := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}
	return months[monthNumber - 1]
}

func GetMonthName2(monthNumber int) (month string) {
	switch monthNumber{
	case 1:
		month = "Enero"
	case 2:
		month = "Febrero"
	case 3:
		month = "Marzo"
	case 4:
		month = "Abril"
	case 5: 
		month = "Mayo"
	case 6:
		month = "Junio"
	case 7:
		month = "Julio"
	case 8:
		month = "Agosto"
	case 9:
		month = "Septiembre"
	case 10:
		month = "Octubre"
	case 11:
		month = "Noviembre"
	case 12:
		month = "Diciembre"
	default:
		month = "Mes invalido"
	}
	return
}