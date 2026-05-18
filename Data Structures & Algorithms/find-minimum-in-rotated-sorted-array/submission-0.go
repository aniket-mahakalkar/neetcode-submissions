func findMin(nums []int) int {


	s := 0

	e := len(nums)-1



	for s < e {

		mid := s + (e-s)/2

		if nums[mid] > nums[e] {

			s = mid +1
			
		}else {

			e = mid

		}



}

return nums[s]

}
