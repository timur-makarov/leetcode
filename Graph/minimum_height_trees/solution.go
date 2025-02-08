package minimum_height_trees

func RunSolution(n int, edges [][]int) []int {
	// WORKS, BUT TAKES TOO LONG FOR LARGE GRAPHS

	//graph := make([][]int, n)
	//visited := make([]int, n)
	//
	//for _, edge := range edges {
	//	graph[edge[1]] = append(graph[edge[1]], edge[0])
	//	graph[edge[0]] = append(graph[edge[0]], edge[1])
	//}

	//var dfs func(int, int, int, int) int
	//dfs = func(rootIdx, prevIdx, vertexIdx, count int) int {
	//	count++
	//	for _, j := range graph[vertexIdx] {
	//		if j != prevIdx {
	//			visited[rootIdx] = max(dfs(rootIdx, vertexIdx, j, count), visited[rootIdx])
	//		}
	//	}
	//	return count
	//}
	//
	//for i := range n {
	//	dfs(i, i, i, 0)
	//}
	//
	//result := make(map[int][]int)
	//minVal := math.MaxInt
	//
	//for i, val := range visited {
	//	if val <= minVal {
	//		minVal = val
	//		result[minVal] = append(result[minVal], i)
	//	}
	//}
	//
	//return result[minVal]

	if n == 1 {
		return []int{0}
	}

	graph := make([][]int, n)
	inDegree := make([]int, n)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
		inDegree[edge[0]]++
		inDegree[edge[1]]++
	}

	var leaves []int

	for i, degree := range inDegree {
		if degree == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		size := len(leaves)
		n -= size

		for _, leafIdx := range leaves {
			inDegree[leafIdx]--

			for _, nextVertexIdx := range graph[leafIdx] {
				inDegree[nextVertexIdx]--

				if inDegree[nextVertexIdx] == 1 {
					leaves = append(leaves, nextVertexIdx)
				}
			}
		}

		leaves = leaves[size:]
	}

	return leaves
}
