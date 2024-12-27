package fruits_into_baskets

func RunSolution(fruits []int) int {
	wStart := 0
	maxLength := 0
	hm := make(map[int]int)

	for wEnd := 0; wEnd < len(fruits); wEnd++ {
		hm[fruits[wEnd]]++

		for len(hm) > 2 {
			hm[fruits[wStart]]--
			if hm[fruits[wStart]] == 0 {
				delete(hm, fruits[wStart])
			}
			wStart++
		}

		maxLength = max(maxLength, wEnd-wStart+1)
	}

	return maxLength
}
