func missingNumber(nums []int) int {

	mis := len(nums)


	for i:= 0; i < len(nums) ; i++ {
		mis ^= i ^ nums[i]
	}

	return mis

}
