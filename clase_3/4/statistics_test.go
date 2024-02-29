package statistics
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMinValue(t *testing.T) {
	//arrange
	values := []int{3,2,7,9,1}
	expected_result := 1
	//act
	result := minValue(values...)
	//assert
	assert.Equal(t, expected_result, result)
}

func TestMaxValue(t *testing.T) {
	//arrange
	values := []int{3,2,7,9,1}
	expected_result := 9
	//act
	result := maxValue(values...)
	//assert
	assert.Equal(t, expected_result, result)
}


func TestAverageValue(t *testing.T) {
	//arrange
	values := []int{3,2,7,9,1,2}
	expected_result := 4
	//act
	result := averageValue(values...)
	//assert
	assert.Equal(t, expected_result, result)
}