package task_scheduler

func RunSolution(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}

	freqs := make([]int, 26)
	for _, task := range tasks {
		freqs[task-'A']++
	}

	maxCount := 0
	sameMaxCount := 0

	for _, freq := range freqs {
		if freq > maxCount {
			maxCount = freq
			sameMaxCount = 1
		} else if freq == maxCount {
			sameMaxCount++
		}
	}

	res := (n+1)*(maxCount-1) + sameMaxCount
	if res > len(tasks) {
		return res
	} else {
		return len(tasks)
	}
}
