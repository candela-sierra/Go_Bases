package main

type Animal string
type foodForAnimal func(float64) (float64)
type FoodByAnimal = float64

const (
	dog Animal = "dog"
	cat Animal= "cat"
	tarantula Animal = "tarantula"
	hamster Animal = "hamster"
	foodByCat FoodByAnimal = 5
	foodByDog FoodByAnimal = 10
	foodByHamster FoodByAnimal = 0.250
	foodByTarantula FoodByAnimal = 0.150
)

func main()  {
	
}

func foodFor(animal Animal) (foodFor foodForAnimal, err string) {
	switch animal {
	case cat:
		foodFor = foodForCats
	case dog:
		foodFor = foodForDogs
	case tarantula:
		foodFor = foodForTarantulas
	case hamster:
		foodFor = foodForHamsters
	default:
		err = "Unknown Animal"
	}
	return
}

func foodForCats(amount float64) float64 {
	return amount * foodByCat
}

func foodForDogs(amount float64) float64 {
	return amount * foodByDog
}

func foodForHamsters(amount float64) float64 {
	return amount * foodByHamster
}

func foodForTarantulas(amount float64) float64 {
	return amount * foodByTarantula
}