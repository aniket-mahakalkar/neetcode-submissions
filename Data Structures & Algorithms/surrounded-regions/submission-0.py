class Solution:
    def solve(self, board: List[List[str]]) -> None:


        rows , cols = len(board), len(board[0])



        visited = [[0 for c in range(cols)] for r in range(rows)]


        def dfs(row,col):

            if row < 0 or col < 0 or row >= rows or col >= cols or visited[row][col] == 1 or board[row][col] != 'O':

                return 

            visited[row][col] = 1

            dfs(row+1,col)
            dfs(row-1,col)
            dfs(row,col-1)
            dfs(row,col+1)




        for  r in range(rows):


            if board[r][0] == 'O' and visited[r][0] == 0:

                dfs(r,0)


            if board[r][cols-1] == 'O' and visited[r][cols-1] == 0:

                dfs(r,cols-1)


        for c in range(cols):

            if board[0][c] == 'O' and visited[0][c] == 0:

                dfs(0,c)

            if board[rows-1][c] == 'O' and visited[rows-1][c]  == 0:

                dfs(rows-1,c)


        
        for r in range(rows):

            for c in range(cols):


                if board[r][c] == 'O' and visited[r][c] == 0:

                    board[r][c] = 'X'

        
        