// house robber pattern but start from the back and look right
func mostPoints(questions [][]int) int64 {
    dp := make([]int64, len(questions))
    dp[len(questions)-1] = int64(questions[len(questions)-1][0])
    for i := len(questions)-2; i >= 0; i-- {
        pt := int64(questions[i][0])
        power := questions[i][1]
        solve := pt
        if i+power+1 < len(dp) {solve += dp[i+power+1]}
        skip := dp[i+1]
        dp[i] = max(solve,skip)
    }
    return dp[0]
}

// house robber pattern
// brute force dfs ( TLE )
// func mostPoints(questions [][]int) int64 {
//     var dfs func(i int, points int64) int64
//     dfs = func(i int, points int64) int64 {
//         // base
//         if i >= len(questions) {return points}
        
//         // logic

//         pt := int64(questions[i][0])
//         power := questions[i][1]

//         solve := dfs(i+power+1, points+pt)
//         skip := dfs(i+1, points)
//         return max(solve, skip)
//     }
//     return dfs(0, int64(0))
// }

func max(x,  y int64) int64 {
    if x > y {return x}
    return y
}