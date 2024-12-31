package interval_list_intersections

func RunSolution(firstList [][]int, secondList [][]int) [][]int {
	var result [][]int
	i := 0
	j := 0

	for i < len(firstList) && j < len(secondList) {
		f_o_s := firstList[i][0] >= secondList[j][0] && firstList[i][0] <= secondList[j][1]
		s_o_f := secondList[j][0] >= firstList[i][0] && secondList[j][0] <= firstList[i][1]

		if f_o_s || s_o_f {
			start := max(firstList[i][0], secondList[j][0])
			end := min(firstList[i][1], secondList[j][1])
			result = append(result, []int{start, end})
		}

		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}

	return result
}
