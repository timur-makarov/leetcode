package maximum_frequency_stack

type FreqStack struct {
	NumToFreq  map[int]int
	freqToNums map[int][]int
	maxFreq    int
}

func Constructor() FreqStack {
	return FreqStack{
		freqToNums: make(map[int][]int),
		NumToFreq:  make(map[int]int),
	}
}

func (fq *FreqStack) Push(val int) {
	fq.NumToFreq[val]++
	fq.freqToNums[fq.NumToFreq[val]] = append(fq.freqToNums[fq.NumToFreq[val]], val)
	fq.maxFreq = max(fq.maxFreq, fq.NumToFreq[val])
}

func (fq *FreqStack) Pop() int {
	lastIndex := len(fq.freqToNums[fq.maxFreq]) - 1
	num := fq.freqToNums[fq.maxFreq][lastIndex]
	fq.freqToNums[fq.maxFreq] = fq.freqToNums[fq.maxFreq][0:lastIndex]

	fq.NumToFreq[num]--

	if len(fq.freqToNums[fq.maxFreq]) == 0 {
		fq.maxFreq--
	}

	return num
}

func RunSolution() []int {
	fq := Constructor()
	fq.Push(5)
	fq.Push(7)
	fq.Push(5)
	fq.Push(7)
	fq.Push(4)
	fq.Push(5)

	return []int{fq.Pop(), fq.Pop(), fq.Pop(), fq.Pop()}
}
