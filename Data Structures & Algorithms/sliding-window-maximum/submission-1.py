class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:

        ans = []
        for right in range(len(nums) -k + 1):

            ans.append(max(nums[right:right+k]))



        return ans





