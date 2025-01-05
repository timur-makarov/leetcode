package String

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/String/generalized_abbreviations"
	"testing"
)

func TestGeneralizedAbbreviations(t *testing.T) {
	s := "BAT"
	answer := []string{"BAT", "BA1", "B1T", "B2", "1AT", "1A1", "2T", "3"}
	result := generalized_abbreviations.RunSolution(s)
	assert.Equal(t, answer, result)
}
