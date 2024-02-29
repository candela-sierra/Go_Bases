package animals

type Animal string
type foodForAnimal func(float64) (float64)
type FoodByAnimal = float64

const (
	Dog Animal = "dog"
	Cat Animal= "cat"
	Tarantula Animal = "tarantula"
	Hamster Animal = "hamster"
	foodByCat FoodByAnimal = 5
	foodByDog FoodByAnimal = 10
	foodByHamster FoodByAnimal = 0.250
	foodByTarantula FoodByAnimal = 0.150
)

func FoodFor(animal Animal) (foodFor foodForAnimal, err string) {
	switch animal {
	case Cat:
		foodFor = foodForCats
	case Dog:
		foodFor = foodForDogs
	case Tarantula:
		foodFor = foodForTarantulas
	case Hamster:
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