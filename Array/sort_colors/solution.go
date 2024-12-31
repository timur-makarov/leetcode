package sort_colors

func RunSolution(nums []int) {
	wStart := 0
	wEnd := len(nums) - 1
	i := 0

	for i < wEnd {
		if nums[i] == 0 {
			nums[i], nums[wStart] = nums[wStart], nums[i]
			wStart++
			i++
		} else if nums[i] == 1 {
			i++
		} else {
			nums[i], nums[wEnd] = nums[wEnd], nums[i]
			wEnd--
		}
	}
}
