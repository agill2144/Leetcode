func longestLine(mat [][]int) int {
    m := len(mat)
    n := len(mat[0])

    type node struct {
        diagUpLeft int
        diagUpRight int
        left int
        up int
    }
    maxLine := 0
    dp := make([][]*node, m)
    for i := 0; i < m; i++ {dp[i] = make([]*node,n)}
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if mat[i][j] == 1 {
                dp[i][j] = &node{1,1,1,1}
                if i-1 >= 0 && j-1 >= 0 && dp[i-1][j-1] != nil {
                    dp[i][j].diagUpLeft += dp[i-1][j-1].diagUpLeft
                }
                if i-1 >= 0 && j+1 < n && dp[i-1][j+1] != nil {
                    dp[i][j].diagUpRight += dp[i-1][j+1].diagUpRight
                }
                if i-1 >= 0 && dp[i-1][j] != nil {
                    dp[i][j].up += dp[i-1][j].up
                }
                if j-1 >= 0 && dp[i][j-1] != nil {
                    dp[i][j].left += dp[i][j-1].left
                }
                maxLine = max(maxLine, 
                    max(dp[i][j].left, 
                    max(dp[i][j].up, 
                    max(dp[i][j].diagUpLeft, dp[i][j].diagUpRight))))
            }
        }
    }
    return maxLine
}