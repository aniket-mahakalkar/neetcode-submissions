func longestPalindrome(s string) string {
        n := len(s)
    
    maxLen := 0
    sInd := 0
    eInd := 0
    
    var expand func(left, right int) 
    
    expand = func(left, right int) {
        
        for left >= 0 && right < n && s[left] == s[right] {
            
            if right - left + 1 > maxLen {
                
                maxLen = right - left + 1
                sInd = left
                eInd = right
                
            }
            
            left--
            right++
        }
    }
    
    
    for i := 0; i < n; i++ {
        
        expand(i,i)
        expand(i,i+1)
    }
    
    return s[sInd:eInd+1]
}
