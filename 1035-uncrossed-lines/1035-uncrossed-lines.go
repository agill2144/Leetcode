// DP: Longest Common Subsequence pattern
func maxUncrossedLines(nums1 []int, nums2 []int) int {
    n1 := len(nums1)
    n2 := len(nums2)
    dp := make([][]int,n1+1)
    for i := 0; i < len(dp); i++ {dp[i] = make([]int, n2+1)}
    
    for i := 1; i < len(dp); i++ {
        for j := 1; j < len(dp[0]); j++ {
            one := nums1[i-1]
            two := nums2[j-1]
            if one == two {
                dp[i][j] = 1+dp[i-1][j-1]
            } else {
                dp[i][j] = max(dp[i-1][j], dp[i][j-1])
            }
        }
    }
    return dp[n1][n2]
}

func max(x, y int) int {
    if x > y {return x}
    return y
}
// brute force dfs with backtracking (TLE)
// func maxUncrossedLines(nums1 []int, nums2 []int) int {
//     idxMap := map[int][]int{}
//     for i := 0; i < len(nums2); i++ {
//         idxMap[nums2[i]] = append(idxMap[nums2[i]], i)
//     }
//     connected := make([]bool, len(nums2))
//     ans := 0
    
//     var dfs func(start int, count int)
//     dfs = func(start int, count int) {
//         // base
//         if count > ans {ans = count}
//         if start == len(nums1) {return}
        
//         // logic
//         for i := start; i < len(nums1); i++ {
//             num := nums1[i]
//             for _, idx := range idxMap[num] {
//                 if !connected[idx] && !doesIntersect(connected, idx) {
//                     connected[idx] = true            
//                     dfs(i+1, count+1)
//                     connected[idx] = false
//                 }
//             }
//         }
//     }
//     dfs(0,0)
//     return ans
// }

// func doesIntersect(arr []bool, idx int) bool {
//     for i := idx+1; i < len(arr); i++ {
//         if arr[i] {return true}
//     }
//     return false
// }