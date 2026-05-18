class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:

        ans = 0
        l =0

        n = len(s)
        s = list(s)
        w = []
        for r in range(n):
            
            while s[r] in w:

                w.pop(0)

            w += s[r]

            ans = max(ans, len(w))

        return ans


            






            

            




        