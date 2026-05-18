class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:


        curr = prev = 0


        for i in range(2,len(cost)+1):

            temp = curr 

            curr = min( (cost[i-2] + prev) , curr + cost[i-1])

            prev = temp


        return curr

            


        