/*
    approach: bottom up dp
    time: o(s^2)
    space: o(s)
*/
func minCut(s string) int {
    dp := make([]int, len(s))
    palindromeCache := map[string]bool{}
    dp[0] = 0
    for i := 1; i < len(dp); i++ {dp[i] = math.MaxInt64}
    for i := 1; i < len(s); i++ {
        for j := i; j >= 0; j-- {
            if isPalindrome(s, j, i, palindromeCache) {
                if j == 0 {dp[i] = 0; break}
                dp[i] = min(dp[i], dp[j-1]+1)
            }
        }
    }
    return dp[len(dp)-1]
}
func min(x, y int) int {
    if x < y {return x}
    return y
}
func isPalindrome(s string, left, right int, palindromeCache map[string]bool) bool {
    if val, ok := palindromeCache[string(s[left:right+1])]; ok {return val}
    for left < right {
        if s[left] != s[right] {palindromeCache[string(s[left:right+1])]=false; return false}
        left++
        right--
    }
    palindromeCache[string(s[left:right+1])] = true
    return true
}