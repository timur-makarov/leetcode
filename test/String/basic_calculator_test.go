package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/basic_calculator"
	"testing"
)

func TestBasicCalculator(t *testing.T) {
	s := "1+(4+5+2)-3+(6+8)"
	answer := 23
	result := basic_calculator.RunSolution(s)
	assert.Equal(t, answer, result)
}
