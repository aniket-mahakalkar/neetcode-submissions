from collections import defaultdict
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:

        hash_map = defaultdict(int)
        
        for i in range(len(nums)):

            if target - nums[i] in hash_map:

                return [hash_map[target - nums[i]] , i]
        
            hash_map[nums[i]] = i
        
        return



        

        

            
        