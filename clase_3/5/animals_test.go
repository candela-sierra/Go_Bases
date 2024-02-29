package animals
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestDogs(t *testing.T) {
	//arrange
	var amount float64 = 5
	var expected_result float64 = 50
	//act
	result := foodForDogs(amount)
	//assert
	assert.Equal(t, expected_result, result)
}

func TestCats(t *testing.T) {
	//arrange
	var amount float64 = 2
	var expected_result float64 = 10
	//act
	result := foodForCats(amount)
	//assert
	assert.Equal(t, expected_result, result)
}

func TestTarantulas(t *testing.T) {
	//arrange
	var amount float64 = 2
	var expected_result float64 = 0.3
	//act
	result := foodForTarantulas(amount)
	//assert
	assert.Equal(t, expected_result, result)
}

func TestHamsters(t *testing.T) {
	//arrange
	var amount float64 = 3
	var expected_result float64 = 0.75
	//act
	result := foodForHamsters(amount)
	//assert
	assert.Equal(t, expected_result, result)
}