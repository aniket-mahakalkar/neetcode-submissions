class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        
        n = len(nums)


        prefix_pro = [1]*n
        suffix_pro = [1]*n

        for i in range(1,n):

            prefix_pro[i] = prefix_pro[i-1] * nums[i-1]

        for i in range(n-2,-1,-1):

            suffix_pro[i] = suffix_pro[i+1] * nums[i+1]

        ans = []

        for i in range(n):

            ans.append(prefix_pro[i] * suffix_pro[i])

        
        return ans

        
        