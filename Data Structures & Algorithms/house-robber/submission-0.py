class Solution:
    def rob(self, nums: List[int]) -> int:

        prev2  = 0
        prev = nums[0]



        for i in range(1, len(nums)):


            take = nums[i] + prev2

            nottake = prev


            curr = max(take, nottake)


            prev2 = prev
            prev = curr

        return prev

            



        