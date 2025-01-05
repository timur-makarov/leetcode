package basic_calculator

func RunSolution(s string) int {
	var calc func(i int) (int, int)
	calc = func(i int) (int, int) {
		var result int
		num := 0
		op := 1

		for i < len(s) {
			l := s[i]

			if l > 47 && l < 58 {
				num = num*10 + int(l-'0')
			} else if l == '+' || l == '-' {
				result += num * op
				num, op = 0, 44-int(l)
			} else if l == '(' {
				num, i = calc(i + 1)
			} else if l == ')' {
				result += num * op
				return result, i
			}

			i++
		}

		result += num * op
		return result, i
	}

	result, _ := calc(0)
	return result
}
