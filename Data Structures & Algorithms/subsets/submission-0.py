class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:


        ans = []
        

        def recursive(ind,ds):
            
            if ind == len(nums):

                ans.append(ds[:])

                return

            ds.append(nums[ind])
            recursive(ind+1 , ds)
            ds.pop()
            recursive(ind+1,ds)


        recursive(0,[])

        return ans


