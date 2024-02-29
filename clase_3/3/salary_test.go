package salary
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSalaryByCategory(t *testing.T) {
	t.Run("test category C", func (t *testing.T)  {
		//arrange
		category := C
		minutes := 120
		expected := 2000
		//act
		result := SalaryByCategory(category, minutes)
		//assert
		assert.Equal(t, expected, result)
	})

	t.Run("test category B", func (t *testing.T)  {
		//arrange
		category := B
		minutes := 120
		expected := 3600
		//act
		result := SalaryByCategory(category, minutes)
		//assert
		assert.Equal(t, expected, result)	
	})

	t.Run("test category A", func (t *testing.T)  {
		//arrange
		category := A
		minutes := 120
		expected := 9000
		//act
		result := SalaryByCategory(category, minutes)
		//assert
		assert.Equal(t, expected, result)
	})
}