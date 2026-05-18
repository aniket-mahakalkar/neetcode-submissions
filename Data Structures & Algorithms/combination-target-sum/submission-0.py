class Solution:
    def combinationSum(self, nums: List[int], target: int) -> List[List[int]]:

        ans = []
        

        def recursive(ind,target, ds):


            if ind == len(nums):
                
                if target == 0:

                    ans.append(ds[:])

                return

            
            if nums[ind] <= target:

                ds.append(nums[ind])
                recursive(ind,target - nums[ind] , ds)
                ds.pop()

            recursive(ind+1 , target, ds)

        recursive(0,target,[])

        return ans



                


