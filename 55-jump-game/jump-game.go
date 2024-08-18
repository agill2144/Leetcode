func canJump(nums []int) bool {
    var dfs func(start int) bool
    memo := map[int]*bool{}
    dfs = func(start int) bool {
        // base
        if start >= len(nums)-1 {return true}

        // logic
        val2, ok2 := memo[start]
        if ok2 {return *val2}
        if nums[start] == 0 { return false }

        for k := nums[start]; k >= 1 ; k-- {
            val, ok := memo[start+k]
            if !ok {
                memo[start+k] = toPtrBool(dfs(start+k))
                if *memo[start+k] {return true} 
            } else {
                return *val
            }
        }
        return false
    }
    return dfs(0)
}
func toPtrBool(b bool) *bool {return &b }
// func canJump(nums []int) bool {
//     farthestIdx := 0
//     i := 0
//     n := len(nums)
//     for i < n && i <= farthestIdx {
//         farthestIdx = max(farthestIdx, i+nums[i])
//         i++
//     }
//     return farthestIdx >= n-1
// }