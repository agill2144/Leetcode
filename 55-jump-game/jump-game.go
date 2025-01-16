// jump in windows
func canJump(nums []int) bool {
    right := 0
    i := 0
    for i <= right {
        right = max(right, i+nums[i])
        if right >= len(nums)-1 {return true}
        i++
    }
    return false
}
// recursively try all possible jump sizes
// be greedy and start with highest possible jump from each position
// numberOfOptionsPerNode^totalNumberOfOptions
// numberOfOptionsPerNode = k
// totalNumberOfOptions = n
// tc = o(k^n)
// sc = o(n)
// func canJump(nums []int) bool {
//     var dfs func(idx int) bool
//     dfs = func(idx int)bool {
//         // base
//         if idx >= len(nums)-1 {return true}
//         // logic
//         jumpSizes := nums[idx]
//         for i := jumpSizes; i >= 1; i-- {
//             if dfs(idx+i) {return true}
//         }
//         return false
//     }
//     return dfs(0)
// }

