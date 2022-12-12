// time: o(nk)
// space: o(n)
func maxSumAfterPartitioning(arr []int, k int) int {
    n := len(arr)
    dp := make([]int, n)
    dp[0] = arr[0]
    
    for i := 1; i < n; i++ {
        
        maxForThisSubProblem := math.MinInt64
        maxInPart := math.MinInt64
        for j := i; j > i-k && j >= 0 ; j-- {
            if arr[j] > maxInPart { maxInPart = arr[j] }
            partSize := i-j+1
            partitionSum := maxInPart * partSize
            if j > 0 {
                partitionSum += dp[j-1]
            }
            maxForThisSubProblem = max(partitionSum, maxForThisSubProblem)
        }
        
        dp[i] = maxForThisSubProblem
    }
    
    return dp[len(dp)-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}