func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}


func canJump(nums []int) bool {

	var maxInd int


	for i:=0; i < len(nums); i++{

		if i > maxInd {
			return false
		}

		maxInd = max(maxInd, nums[i] + i)


	}
	
	return true
}
