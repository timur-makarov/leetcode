package longest_substring_without_repeating_characters

func RunSolution(s string) int {
	maxLength := 0
	wStart := 0
	hm := map[uint8]int{}

	for wEnd := 0; wEnd < len(s); wEnd++ {
		i, ok := hm[s[wEnd]]

		if ok {
			wStart = max(wStart, i+1)
		}

		hm[s[wEnd]] = wEnd
		maxLength = max(maxLength, wEnd-wStart+1)
	}

	return maxLength
}
