func canJump(nums []int) bool {
    dp := make([]bool, len(nums))
    dp[len(dp)-1] = true
    for i := len(dp)-2; i >= 0; i-- {
        numJumps := nums[i]
        for j := numJumps; j >= 1; j-- {
            if i+j < len(nums) && dp[i+j] {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[0]
}
// func canJump(nums []int) bool {
//     memo := map[int]bool{}
//     var dfs func(start int) bool
//     dfs = func(start int) bool {
//         // base
//         if start >= len(nums)-1 {
//             if start == len(nums)-1{return true}
//             return false
//         }
        
//         // logic
//         if _, ok := memo[start]; !ok {
//             numJumps := nums[start]
//             for j := numJumps; j >= 1; j-- {
//                 if ok := dfs(start+j); ok {memo[start]=true; return true}
//             }
//             memo[start] = false
//         }
//         return memo[start]
//     }
//     return dfs(0)
// }