package main
import "fmt"

func main() {
	word := "palabra"
	fmt.Println("Cantidad de letras: ", len(word))
	SpellWord(word)
}

func SpellWord(word string) {
	for _, letter := range word {
		fmt.Printf("%c ", letter)
	}
	fmt.Println("")
}