package Heap

import (
	"github.com/stretchr/testify/assert"
	"github.com/timur-makarov/leetcode/Heap/kth_smallest_element_in_a_sorted_matrix"
	"testing"
)

func TestKthSmallestElementinaSortedMatrix(t *testing.T) {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	k := 8
	answer := 13
	result := kth_smallest_element_in_a_sorted_matrix.RunSolution(matrix, k)
	assert.Equal(t, answer, result)
}
