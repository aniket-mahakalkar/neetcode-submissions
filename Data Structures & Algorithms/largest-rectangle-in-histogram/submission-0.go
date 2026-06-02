func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {

		stack := []int{}
	maxArea := 0

	heights = append(heights, 0)

	for i, h := range heights {

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {

			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			
			if len(stack) == 0 {
				maxArea = max(maxArea, height*i)
			}else {
				maxArea = max(maxArea, height*(i-stack[len(stack)-1]-1))
			}
	
		}

		stack = append(stack, i)

	}

	return maxArea

}
