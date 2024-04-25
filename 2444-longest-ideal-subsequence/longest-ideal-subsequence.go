func longestIdealString(s string, k int) int {
    if len(s) <= 1 {return len(s)}
    dp := make([]int, 26)
    
    longest := 0
    for i := 0; i < len(s); i++ {
        char := s[i]
        charIdx := int(char-'a')


        // k choices on left and k choices on right
        leftEnd := charIdx-k
        rightEnd := charIdx+k
        best := dp[charIdx]
        for j := charIdx-1; j >= 0 && j >= leftEnd; j-- {
            best = max(best, dp[j])
        }
        for j := charIdx+1; j < 26 && j <= rightEnd; j++ {
            best = max(best, dp[j])
        }
        dp[charIdx] = best+1
        longest = max(longest, dp[charIdx])
    }
    return longest
}


// choose / not choose brute force : TLE
// func longestIdealString(s string, k int) int {
//     longest := math.MinInt64
//     var dfs func(ptr int, path string)
//     dfs = func(ptr int, path string) {
//         // base
//         longest = max(longest, len(path))
//         if ptr == len(s) {return}

//         // logic
//         // not-choose
//         dfs(ptr+1, path)
//         // choose
//         if len(path) == 0 {
//             dfs(ptr+1, path+string(s[ptr]))
//         } else {
//             // can we choose ?
//             lastChar := path[len(path)-1]
//             currChar := s[ptr]
//             if abs(int(lastChar)-int(currChar)) <= k {
//                 // yes, we can
//                 dfs(ptr+1, path+string(currChar))
//             } else {
//                 // no we cannot
//                 dfs(ptr+1, path)
//             }
//         }
//     }
//     dfs(0, "")
//     return longest
// }

func abs(x int) int {
    if x < 0 { return x * -1 }
    return x
}