

func exist(board [][]byte, word string) bool {



    rows := len(board)

    cols := len(board[0])


    var dfs func(int, int, int) bool


    dfs = func(r,c, ind int) bool {

        if ind == len(word){

            return true
        }

        if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[ind]{

            return false
        }


        temp := board[r][c]

        board[r][c] = '#'
        found := dfs(r+1,c,ind+1, ) || dfs(r-1,c,ind+1) || dfs(r,c+1,ind+1) || dfs(r,c-1,ind+1)

        board[r][c] = temp


        return found

    }


    for r := 0; r < rows; r++ {

        for c := 0; c < cols; c++ {


            if board[r][c] == word[0] {
                
                if dfs(r,c,0) {
                    return true
                }
            }
        }
    }

    return false

	


}
