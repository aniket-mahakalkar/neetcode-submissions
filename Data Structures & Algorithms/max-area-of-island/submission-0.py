class Solution:
    def maxAreaOfIsland(self, grid: List[List[int]]) -> int:



        rows , cols = len(grid), len(grid[0])

        
        def traverse_island(r,c):

            if r < 0 or c < 0 or r >= rows or c >= cols or grid[r][c] != 1:

                return 0

            
                
            grid[r][c] = 0

            return 1 + traverse_island(r+1,c) + traverse_island(r-1,c) + traverse_island(r,c-1) +traverse_island(r,c+1)
                
                
                
                

        
            


        max_area = 0
        for r in range(rows):

            for c in range(cols):


                if grid[r][c] == 1:

                    

                    max_area = max(max_area ,   traverse_island(r, c))

        return max_area


        