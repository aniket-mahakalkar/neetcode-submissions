from collections import Counter

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:


        cnt1 = Counter(s)
        cnt2 = Counter(t)

        if len(s) != len(t):

            return False

        for key, val in cnt1.items():

            if cnt1[key] != cnt2[key]:

                return False

        return True


        