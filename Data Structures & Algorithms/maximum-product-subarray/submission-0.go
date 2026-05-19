func max(a, b int) int {
	
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {

	ans := nums[0]
	maxPro := nums[0]
	minPro := nums[0]
    
	for i:=1; i < len(nums); i++ {

		prevMax := maxPro
		prevMin := minPro


		maxPro = max(nums[i], max(prevMax*nums[i], prevMin*nums[i]))
		minPro = min(nums[i], min(prevMax*nums[i], prevMin*nums[i]))


		ans = max(ans, maxPro)
	}

	return ans
}
