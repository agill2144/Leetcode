/*
    approach: bottom up dp
    time: o(s^3)
        - o(s) for looping over s
        - o(s) for each char, we go back all the way to 0 idx ; worst case
        - o(s) for palindrome check
    space: o(s)
*/
func minCut(s string) int {
    dp := make([]int, len(s))
    dp[0] = 0
    for i := 1; i < len(dp); i++ {dp[i] = math.MaxInt64}
    for i := 1; i < len(s); i++ {
        for j := i; j >= 0; j-- {
            if isPalindrome(s, j, i) {
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
func isPalindrome(s string, left, right int) bool {
    for left < right {
        if s[left] != s[right] {return false}
        left++
        right--
    }
    
    return true
}