// approach: bottom up DP
func findBall(grid [][]int) []int {
    m := len(grid)
    n := len(grid[0])
    
    dp := make([][]int, m)
    for i := 0; i < m; i++ {dp[i] = make([]int, n)}
    out := make([]int, n)
    
    for i := m-1; i >= 0; i-- {
        for j := 0; j < n; j++ {
            currDir := grid[i][j]
            r := i+1
            c := j+grid[i][j]
            if (c < 0 || c == n || (currDir == 1 && grid[i][c] == -1) || (currDir == -1 && grid[i][c] == 1)) {
                dp[i][j] = -1
            } else {
                result := c
                if r < m {result = dp[r][c]}
                dp[i][j] = result
            }
            if i == 0 {
                out[j] = dp[i][j]
            }
        }
    }
    return out
}

// approach: simulate and identify cols we fall down from
// func findBall(grid [][]int) []int {
//     m := len(grid)
//     n := len(grid[0])
//     // 1  -> \
//     // -1 -> /
    
//     var follow func(r, c int) (bool, int)
//     follow = func(r, c int) (bool, int) {
//         for r >= 0 && r < m && c >= 0 && c < n {
//             currDir := grid[r][c]
//             if currDir == 1 {
//                 if (c+1 < n && grid[r][c+1] == -1) || c+1 == n {return false, -1}
//                 r++; c++;
//             } else if currDir == -1 {
//                 if (c-1 >= 0 && grid[r][c-1] == 1) || c-1 < 0 {return false,-1}
//                 r++; c--
//             }
//         }
//         return true,c
//     }

//     out := make([]int, n)
//     for i := 0; i < n ; i++ {
//         if ok, cIdx := follow(0, i); ok {
//             out[i] = cIdx
//         } else if !ok {
//             out[i] = -1
//         }
//     }
//     return out
// }

