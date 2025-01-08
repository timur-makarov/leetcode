package complement_of_base_10_integer

import (
	"math"
)

func RunSolution(n int) int {
	if n == 0 {
		return 7
	}

	bitsCount := 0
	shiftingNum := n

	for shiftingNum > 0 {
		bitsCount++
		shiftingNum >>= 1
	}

	// 16 = 10000 15 = 1111
	allOnes := int(math.Pow(2, float64(bitsCount))) - 1

	return n ^ allOnes
}
