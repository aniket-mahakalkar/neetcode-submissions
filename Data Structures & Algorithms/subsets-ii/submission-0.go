import (
	"slices"
)

func bt(start int, nums *[]int, curr []int, ans *[][]int) {

	temp := make([]int, len(curr))
	copy(temp, curr)

	*ans = append(*ans, temp)


	for i := start; i < len(*nums); i++ {

		// Skip duplicates
		if i > start && (*nums)[i] == (*nums)[i-1] {
			continue
		}

		curr = append(curr, (*nums)[i])

		bt(i+1, nums, curr, ans)

		curr = curr[:len(curr)-1]
	}
}




func subsetsWithDup(nums []int) [][]int {

slices.Sort(nums)
	ans := [][]int{}

	bt(0, &nums, []int{}, &ans)

	return ans

}
