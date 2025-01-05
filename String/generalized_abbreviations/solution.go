package generalized_abbreviations

import "fmt"

type Abbreviation struct {
	word       string
	index      int
	holesCount int
}

func RunSolution(s string) []string {
	var result []string
	var queue []Abbreviation
	queue = append(queue, Abbreviation{word: ""})

	for len(queue) > 0 {
		abw := queue[0]
		queue = queue[1:]

		if abw.index == len(s) {
			if abw.holesCount != 0 {
				abw.word += fmt.Sprint(abw.holesCount)
			}

			result = append(result, abw.word)
		} else {
			queue = append(queue, Abbreviation{word: abw.word, index: abw.index + 1, holesCount: abw.holesCount + 1})

			if abw.holesCount != 0 {
				abw.word += fmt.Sprint(abw.holesCount)
			}

			abw.word += string(s[abw.index])

			queue = append(queue, Abbreviation{word: abw.word, index: abw.index + 1})
		}
	}

	return result
}
