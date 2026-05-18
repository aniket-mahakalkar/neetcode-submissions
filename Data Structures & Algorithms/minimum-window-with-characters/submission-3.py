from collections import Counter
class Solution:
    def minWindow(self, s: str, t: str) -> str:

        n = len(s)
        l = 0
        ans = ''
        cnt_t = Counter(t)
        have, need = 0 , len(cnt_t)
        
        hash_map = {}

        for i in range(n):
    
            if s[i] in t:
                
                hash_map[s[i]] = hash_map.get(s[i] , 0 ) + 1
                
                if hash_map[s[i]] == cnt_t[s[i]]:
                    
                    have +=1
                    
                
            while have == need:
                
                # print(s[l:i+1])

                
                if (ans == '')  or (i-l+1 < len(ans)):
                    
                    ans = s[l:i+1]
                
                
                
                if s[l] in t:
                    hash_map[s[l]] -= 1
                    
                    if hash_map[s[l]] < cnt_t[s[l]]:
                        
                        have -= 1
                    
                l +=1



        return ans