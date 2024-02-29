package average

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	//arrange
	grades := []int{10, 8, 8, 10}
	expected := 9
	//act
	result := Average(grades...)
	//assert
	assert.Equal(t, expected, result)
}
