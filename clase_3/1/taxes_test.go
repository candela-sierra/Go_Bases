package taxes
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTaxesBelow50000(t *testing.T) {
	//arrange
	salary := 30000
	expected := 0
	//act
	result := Taxes(salary)
	//assert
	assert.Equal(t, expected, result)
}

func TestTaxesOver50000(t *testing.T) {
	//arrange
	salary := 70000
	expected := 11900
	//act
	result := Taxes(salary)
	//assert
	assert.Equal(t, expected, result)
}

func TestTaxesOver150000(t *testing.T) {
	//arrange
	salary := 200000
	expected := 54000
	//act
	result := Taxes(salary)
	//assert
	assert.Equal(t, expected, result)
}