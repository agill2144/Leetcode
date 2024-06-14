func maxUncrossedLines(nums1 []int, nums2 []int) int {
    n1 := len(nums1)
    n2 := len(nums2)
    dp := make([][]int, n2+1)
    for i := 0 ; i < len(dp); i++ {
        dp[i] = make([]int, n1+1)
    }

    for i := 1; i < len(dp); i++ { // nums2
        for j := 1; j < len(dp[i]); j++ { // nums1
            n1Val := nums1[j-1]
            n2Val := nums2[i-1]
            if n1Val == n2Val {
                dp[i][j] = 1 + dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[n2][n1]
}