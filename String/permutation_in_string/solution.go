package permutation_in_string

func RunSolution(s1, s2 string) bool {
	wStart := 0
	maxLength := 0
	cf := map[uint8]int{}

	for i := 0; i < len(s1); i++ {
		cf[s1[i]]++
	}

	for wEnd := 0; wEnd < len(s2); wEnd++ {
		if _, ok := cf[s2[wEnd]]; ok {
			cf[s2[wEnd]]--
			if cf[s2[wEnd]] == 0 {
				maxLength++
			}
		}

		if maxLength == len(cf) {
			return true
		}

		if wEnd >= len(s1)-1 {
			if _, ok := cf[s2[wStart]]; ok {
				if cf[s2[wStart]] == 0 {
					maxLength--
				}
				cf[s2[wStart]]++
			}
			wStart++
		}
	}

	return false
}
