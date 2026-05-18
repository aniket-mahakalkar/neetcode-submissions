# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def goodNodes(self, root: TreeNode) -> int:

        
        def dfs(node,curr_max):


            if not node:

                return 0
            cnt = 0
            if node.val >= curr_max:

                cnt = 1

                curr_max = node.val

            cnt += dfs(node.left, curr_max)
            cnt += dfs(node.right , curr_max)


            return cnt

        return dfs(root,root.val)




        




        