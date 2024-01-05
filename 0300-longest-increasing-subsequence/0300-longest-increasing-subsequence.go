func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    for i := 0; i < len(dp); i++ {
        dp[i] = 1
    }
    longest := 1
    for i := 1; i < len(nums); i++ {
        curr := nums[i]
        for j := i-1; j >= 0; j-- {
            if curr > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        longest = max(longest, dp[i]) 
    }
    return longest
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
// brute force ; backtracking
// form all possible combinations
// tc = exponential
// func lengthOfLIS(nums []int) int {
//     largest := 0
//     var dfs func(start int, path []int)
//     dfs = func(start int,  path []int) {
//         // base
//         if len(path) > largest {largest = len(path)}
        
//         // logic
//         for i := start; i < len(nums); i++ {
//             if len(path) == 0 || nums[i] > path[len(path)-1] {
//                 path = append(path, nums[i])
//                 dfs(i+1, path)
//                 path = path[:len(path)-1]
//             }
//         }
//     }
//     dfs(0, nil)
//     return largest
// }