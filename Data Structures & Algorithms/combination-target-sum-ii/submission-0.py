class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:

        candidates.sort()


        def backtrack(start,curr,target):

            if target == 0:

                ans.append(curr[:])
                return 
            if target < 0:

                return 

            prev = -1


            for i in range(start, len(candidates)):


                if candidates[i] == prev:

                    continue

                if candidates[i] > target :

                    return 

                curr.append(candidates[i])
                backtrack(i+1,curr,target - candidates[i])
                
                curr.pop()

                prev = candidates[i]

        ans = []
        backtrack(0,[],target)


        return ans

        









        