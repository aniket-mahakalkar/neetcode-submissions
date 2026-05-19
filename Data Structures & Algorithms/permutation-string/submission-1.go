func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {

		return false
	}

	cnt_s1 := [26]int{}
	window := [26]int{}


	for i := 0; i < len(s1); i++ {

		cnt_s1[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	if cnt_s1 == window {

		return true
	}

	left := 0

	for j := len(s1); j < len(s2);j++ {

		window[s2[j]-'a']++
		window[s2[left]-'a']--


		left++


			if cnt_s1 == window {

		return true
	}


	}



	return false




	


}
