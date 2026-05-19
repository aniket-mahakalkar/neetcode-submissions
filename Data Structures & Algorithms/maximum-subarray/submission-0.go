

func max(a,b int) int {

	if a > b {

		return a
	}

	return b
}

func maxSubArray(nums []int) int {



	ans := nums[0]
	currSum := nums[0]

	for i := 1; i < len(nums); i++ {

		currSum = max(nums[i], currSum + nums[i])

		ans = max(ans, currSum)

		

	}


	return ans

	
    
}
