class Solution:
    def isValid(self, s: str) -> bool:


        b = {'(':')', '[':']','{':'}'}

        stack = []
        for x in s :

            if x in b:
                
                stack.append(x)

            else:

                if not stack or b[stack.pop()] != x:

                    return False

        return not stack






        