package Array

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Array/search_in_rotated_sorted_array"
	"log"
	"testing"
)

func TestSearchinRotatedSortedArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	answer := -1
	result := search_in_rotated_sorted_array.RunSolution(nums, target)
	assert.Equal(t, answer, result)

	log.Println("asfdhjklsGFSDKJFHJSDFKLGHJSDLKJGHFJSDHK")

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 6
	answer = 2
	result = search_in_rotated_sorted_array.RunSolution(nums, target)
	assert.Equal(t, answer, result)
	log.Println("asfdhjklsGFSDKJFHJSDFKLGHJSDLKJGHFJSDHK")

	nums = []int{4, 5, 6, 7, 0, 1, 2}
	target = 0
	answer = 4
	result = search_in_rotated_sorted_array.RunSolution(nums, target)
	assert.Equal(t, answer, result)
}
