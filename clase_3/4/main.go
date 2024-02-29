package main

type statisticOperation func(values ...int) (int)
type statisticOperator string
const (
	minimum statisticOperator = "minimum"
	average statisticOperator = "average"
	maximum statisticOperator = "maximum"
)

func main()  {
	
}

func operation(operator statisticOperator) (operation statisticOperation, err string)  {
	switch operator {
	case minimum:
		operation = minValue
	case maximum:
		operation = maxValue
	case average:
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