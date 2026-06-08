func solve(nums []int) int{
    
    var prev1, prev2 int
    
    for _, val := range nums {
        temp := prev1
        curr  := max(prev2 + val, prev1)
        prev2 = temp
        prev1 = curr
        
    }
    
    return prev1
    
}

func rob(nums []int) int {

	if len(nums) < 2 {
		return nums[0]
	}
    max2 := solve(nums[1:])
    max1 := solve(nums[:len(nums)-1])  

	return max(max1, max2)
}
