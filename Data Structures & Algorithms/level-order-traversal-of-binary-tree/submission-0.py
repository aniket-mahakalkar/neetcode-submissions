# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

from collections import deque

class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:

        if not root:

            return []

        queue = deque()

        ans = []

        queue.append(root)
        
        while queue:

            level_size = len(queue)

            curr_list = []

            for i in range(level_size):

                curr = queue.popleft()

                curr_list.append(curr.val)

                if curr.left:

                    queue.append(curr.left)

                if curr.right:

                    queue.append(curr.right)

                
            ans.append(curr_list)

        return ans


            



    


        
        