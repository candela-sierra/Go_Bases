package statistics

type statisticOperation func(values ...int) (int)
type statisticOperator string
const (
	Minimum statisticOperator = "minimum"
	Average statisticOperator = "average"
	Maximum statisticOperator = "maximum"
)

func main()  {
	
}

func Operation(operator statisticOperator) (operation statisticOperation, err string)  {
	switch operator {
	case Minimum:
		operation = minValue
	case Maximum:
		operation = maxValue
	case Average:
		operation = averageValue
	default:
		err = "Operacion no definida"
	}
	return 
}

func minValue(values ...int) (minValue int){
	minValue = values[0]
	for _, value := range values {
		if(value < minValue) {
			minValue = value
		}
	}
	return
}

func maxValue(values ...int) (maxValue int){
	for _, value := range values {
		if(value > maxValue) {
			maxValue = value
		}
	}
	return
}

func averageValue(values ...int) int{
	return sum(values)/len(values)
}

func sum(values []int) (result int) {
	for _, value := range values {
		if(value > 0){
			result += value
		}
	}
	return
}