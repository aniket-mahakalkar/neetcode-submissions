func countSubstrings(s string) int {
	n := len(s) 

	cnt := 0

	var expand func(left, right int)  


	expand = func(left, right int)  {
		

		for left >= 0 && right < n && s[left] == s[right] {
			cnt++
			left--
			right++
		}


	}

	for i := 0; i < n; i++ {

		expand(i, i)
		expand(i, i+1)
		
	}
	return cnt    
}
