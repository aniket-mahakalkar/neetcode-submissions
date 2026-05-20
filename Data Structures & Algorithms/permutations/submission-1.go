


func backtrack(curr []int, nums []int, ans *[][]int, visit map[int]bool){

    if len(curr) == len(nums){

        temp := make([]int,len(curr))

        copy(temp, curr)

        *ans = append(*ans, temp)
    }

    for i :=0; i < len(nums); i++ {


        if visit[i] {

            continue

        }

        visit[i] = true

        curr = append(curr, nums[i])

        backtrack(curr,nums,ans,visit)
        curr = curr[:len(curr)-1]
        visit[i] =false
        
    }

}


func permute(nums []int) [][]int {



    visit := make(map[int]bool)

    ans := [][]int{}

    backtrack([]int{}, nums, &ans, visit)

    return ans

}
