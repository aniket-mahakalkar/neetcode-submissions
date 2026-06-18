func longestCommonSubsequence(text1 string, text2 string) int {

    rows := len(text1)
    cols := len(text2)

    dp := make([][]int, rows+1)

    for i := 0; i <= rows; i++ {
        dp[i] = make([]int, cols+1)
    }

    for row := 1; row <= rows; row++ {

        for col := 1; col <= cols; col++ {

            if text1[row-1] == text2[col-1] {
                dp[row][col] = 1 + dp[row-1][col-1]
            } else {
                dp[row][col] = max(dp[row-1][col], dp[row][col-1])
            }
        }
    }

    return dp[rows][cols]
}