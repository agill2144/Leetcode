func numDecodings(s string) int {
    if len(s) < 1 {return 0}
    dp := make([]int, len(s))
    dp[0] = 1
    if s[0] == '0' {
        dp[0] = 0
    }

    for i := 1; i < len(s); i++ {
        for j := i; j >= 0; j-- {
            if i-j+1 > 2 {break}
            subStr := s[j:i+1]
            if subStr[0] == '0' {continue}
            subStrInt, _ := strconv.Atoi(subStr)
            if subStrInt > 26 {break}
            curr := 1
            if j-1 >= 0 {curr = dp[j-1]}
            dp[i] += curr
        }
    }
    return dp[len(dp)-1]
}