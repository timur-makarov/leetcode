package main

import (
	"log"
)

func main() {
	//sorts.RunSortFunctions()

	log.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	var result []string

	var dfs func(int, int, string)
	dfs = func(open, close int, s string) {
		if len(s) == n*2 {
			result = append(result, s)
			return
		}

		if open < n {
			dfs(open+1, close, s+"(")
		}

		if close < open {
			dfs(open, close+1, s+")")
		}
	}

	dfs(0, 0, "")
	return result
}
