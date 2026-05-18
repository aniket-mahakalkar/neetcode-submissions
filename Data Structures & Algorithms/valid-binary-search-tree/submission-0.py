# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:


        def dfs_check(node,low,high):

            if not node:

                return True

            if not ( low < node.val < high):

                return False

            
            return dfs_check(node.left,low,node.val) and dfs_check(node.right,node.val,high)


        return dfs_check(root,float('-inf'), float('inf'))


            

            

            

            
        