/*
    approach: brute force
    - nested iterations
    time = o(n^2)
    space = o(1)
    
    approach: prefix / suffix / running Sum
    - first pass to get total
    - second pass to get left and right sum
        - if we have a runningSum, and we have total sum
        - to get the leftSum ( we already have it , runningSum )
        - to get the rightSum, remove leftSum from total and add ith val
    
    time = o(2n) ; o(n)
    space = o(1)
*/
func maximumSumScore(nums []int) int64 {
    var total int64
    for i := 0; i < len(nums); i++ {
        total += int64(nums[i])
    }
    var ans int64  = math.MinInt64
    var prefixSum int64
    for i := 0; i < len(nums); i++ {
        prefixSum += int64(nums[i])
        suffixSum := (total-prefixSum)+int64(nums[i])
        ans = max(ans, max(prefixSum, suffixSum))
    }  
    return ans
}

func max(x, y int64) int64 {
    if x > y {return x}
    return y
}