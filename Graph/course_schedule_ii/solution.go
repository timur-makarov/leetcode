package course_schedule_ii

func RunSolution(numCourses int, prerequisites [][]int) []int {
	var result []int
	graph := make([][]int, numCourses)
	visited := make([]int, numCourses)

	for _, prereq := range prerequisites {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	var dfs func(int) bool
	dfs = func(vertexIdx int) bool {
		if visited[vertexIdx] == -1 {
			return false
		}
		if visited[vertexIdx] == 1 {
			return true
		}

		visited[vertexIdx] = -1
		for _, j := range graph[vertexIdx] {
			if !dfs(j) {
				return false
			}
		}
		visited[vertexIdx] = 1
		result = append(result, vertexIdx)

		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return nil
		}
	}

	return result
}
