package main
import "fmt"

func main() {

}

func average(values ...int) int{
	return sum(values)/len(values)
}

func sum(values []int) (result int) {
	for _, value := range values {
		if(value > 0){
			result += value
		} else {
			fmt.Println("Las notas no pueden ser negativas")
		}
	}
	return
}