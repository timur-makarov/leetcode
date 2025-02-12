package main

import "log"

func main() {
	//sorts.RunSortFunctions()

	s := "applepenapple"
	wordDict := []string{"apple", "pen"}
	log.Println(wordBreak(s, wordDict))

	s = "catsandog"
	wordDict = []string{"cats", "dog", "sand", "and", "cat"}
	log.Println(wordBreak(s, wordDict))

	s = "aaaaaaa"
	wordDict = []string{"aaaa", "aaa"}
	log.Println(wordBreak(s, wordDict))
}

func wordBreak(s string, wordDict []string) bool {
	memo := make([]bool, len(s)+1)
	memo[0] = true

	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			if i >= len(word) {
				checkLength := memo[i-len(word)]
				currentSub := s[i-len(word) : i]
				if checkLength && currentSub == word {
					memo[i] = true
				}
			}
		}
	}

	return memo[len(s)]
}
