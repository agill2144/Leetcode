// func canPartition(nums []int) bool {
//     totalSum := 0
//     for i := 0; i < len(nums); i++ {
//         totalSum += nums[i]
//     }
//     if totalSum % 2 != 0 {
//         return false
//     }
//     var helper func(target int, start int) bool
//     helper = func(target int, start int) bool {
//         // base
//         if target == 0 {
//             return true
//         }
//         if target <= 0 || start == len(nums) {
//             return false
//         }
        
//         // logic
//         return helper(target-nums[start], start+1) || helper(target, start+1) 
//     }
    
//     return helper(totalSum/2, 0)
// }


func canPartition(nums []int) bool {
    totalSum := 0
    for i := 0; i < len(nums); i++ {
        totalSum += nums[i]
    }
    if totalSum % 2 != 0 {
        return false
    }
    target := totalSum/2
    m := len(nums)
    n := target
    dp := make([][]bool, m+1)
    for idx, _ := range dp {
        dp[idx] = make([]bool, n+1)
    }
    dp[0][0] = true
    
    for i := 1; i < len(dp); i++ { // value
        for j := 0; j < len(dp[0]); j++ { // target
            val := nums[i-1]
            if val > j {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] || dp[i-1][j-val]
            }
        }
    }
    
    return dp[len(dp)-1][len(dp[0])-1]
}