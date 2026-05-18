from collections import defaultdict
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:

        anagram = defaultdict(list)

        for i in strs:

            sorted_word = ''.join(sorted(i))

            anagram[sorted_word].append(i)

        return list(anagram.values())


        