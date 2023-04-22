// func numDecodings(s string) int {
//     count := 0
//     var dfs func(start int)
//     dfs = func(start int) {
//         // base
//         if start == len(s) {count++; return}
        
//         // logic
//         for i := start; i < len(s); i++ {
//             if s[start] == '0' {continue}
//             subStr := s[start:i+1]
//             subStrInt, _ := strconv.Atoi(subStr)
//             if subStrInt <= 0 || subStrInt > 26 {continue}
//             dfs(i+1)
//         }
//     }
//     dfs(0)
//     return count
// }

func numDecodings(s string) int {
    if s[0] == '0' {return 0}
    
    dp := make([]int, len(s))
    dp[0] = 1
    for i := 1; i < len(s); i++ {
        for j := i; j >= 0; j-- {
            if i-j+1 > 2 {break}
            if s[j] == '0' {continue}
            subStr := string(s[j:i+1])
            subStrInt, _ := strconv.Atoi(subStr)
            if subStrInt > 0 && subStrInt <= 26 {
                oneStepBack := 1
                if j-1 >= 0 {oneStepBack = dp[j-1]}
                dp[i] += oneStepBack
            }
        }
    }
    return dp[len(dp)-1]
}