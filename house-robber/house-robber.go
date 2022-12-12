
// brute force dfs
// time: exponential
// func rob(nums []int) int {
//     max := math.MinInt64
//     var dfs func(start int, profit int)
//     dfs = func(start int, profit int) {
//         // base
//         if profit > max {
//             max = profit
//         }
//         if start >= len(nums) {
//             return
//         }
        
//         // logic
//         dfs(start+2, profit+nums[start])
//         dfs(start+1, profit)
//     }
    
//     dfs(0,0)
//     return max
// }


// top - down using dp matrix
// time: o(n)
// space: o(n)
// func rob(nums []int) int {
//     dp := make([][]int, len(nums))
//     for idx, _ := range dp {
//         dp[idx] = make([]int, 2) // 0 | 1 case
//     }
//     dp[0][0] = 0
//     dp[0][1] = nums[0]
//     for i := 1; i < len(dp); i++ {
//         dp[i][0] = int(math.Max(float64(dp[i-1][0]),float64(dp[i-1][1]))) 
//         dp[i][1] = nums[i] + dp[i-1][0]
//     }
//     if dp[len(dp)-1][0] > dp[len(dp)-1][1] {
//         return dp[len(dp)-1][0]
//     }
//     return dp[len(dp)-1][1]
// }


// bottom up using dp matrix
// time: o(n)
// space: o(n)
// func rob(nums []int) int {
//     dp := make([][]int, len(nums))
//     for idx, _ := range dp {
//         dp[idx] = make([]int, 2) // 0 | 1 case
//     }
//     dp[len(dp)-1][0] = 0
//     dp[len(dp)-1][1] = nums[len(nums)-1]
//     for i := len(dp)-2; i >= 0; i-- {
//         dp[i][0] = int(math.Max(float64(dp[i+1][0]),float64(dp[i+1][1]))) 
//         dp[i][1] = nums[i] + dp[i+1][0]
//     }
//     return int(math.Max(float64(dp[0][1]),float64(dp[0][0])))
// }


// bottom up DP using dp array of fixed size -- can be replaced by 2 vars instead
// time: o(n)
// space: o(1)
// func rob(nums []int) int {
//     dp := make([]int, 2)
//     dp[0] = 0
//     dp[1] = nums[len(nums)-1]
//     for i := len(nums)-2; i >= 0; i-- {
//         choose := dp[1]
//         notChoose := dp[0]
        
//         dp[0] = int(math.Max(float64(choose),float64(notChoose))) 
//         dp[1] = nums[i] + notChoose
//     }
//     return int(math.Max(float64(dp[0]),float64(dp[1])))
// }


// top - down using fixed size ( 2 ) 1D array -- can be replaced with 2 vars
// time: o(n)
// space: o(1)
// func rob(nums []int) int {
//     dp := make([]int, 2)
//     dp[0] = 0
//     dp[1] = nums[0]
//     for i := 1; i < len(nums); i++ {
//         choose := dp[1]
//         notChoose := dp[0]
//         dp[0] = int(math.Max(float64(notChoose),float64(choose))) 
//         dp[1] = nums[i] + notChoose
//     }
//     return int(math.Max(float64(dp[0]),float64(dp[1])))
// }



func rob(nums []int) int {
    dp := make([]int, len(nums)+1)
    dp[1] = nums[0]
    
    /*
        [1,2,3,1]
             i
        [0,1,0,0,0]
    */
    for i := 2; i < len(dp); i++ {
        currHouse := nums[i-1]
        dp[i] = max(currHouse+dp[i-2], dp[i-1])
    }
    return dp[len(dp)-1]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
} 