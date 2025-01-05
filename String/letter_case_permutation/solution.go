package letter_case_permutation

import "fmt"

func RunSolution(s string) []string {
	var result []string
	result = append(result, s)

	for i := range s {
		if isDigit(s[i]) {
			continue
		}

		n := len(result)
		for j := 0; j < n; j++ {
			var newChar byte
			var newStr string

			oldStr := result[j]

			if oldStr[i] >= 'a' && oldStr[i] <= 'z' {
				newChar = oldStr[i] - 32
			} else {
				newChar = oldStr[i] + 32
			}

			newStr = oldStr[:i] + fmt.Sprintf("%c", newChar) + oldStr[i+1:]
			result = append(result, newStr)
		}
	}

	return result
}

func isDigit(l byte) bool {
	return l == 48 || l == 49 || l == 50 || l == 51 || l == 52 || l == 53 || l == 54 || l == 55 || l == 56 || l == 57
}
