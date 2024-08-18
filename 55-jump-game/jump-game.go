// top down memo
func canJump(nums []int) bool {
    var dfs func(start int) bool
    memo := map[int]*bool{}
    dfs = func(start int) bool {
        // base
        if start >= len(nums)-1 {return true}

        // logic
        // if memo[start] != nil {return *memo[start]}
        if nums[start] == 0 { return false }

        for k := nums[start]; k >= 1 ; k-- {
            if memo[start+k] == nil {
                // we dont have an answer for start+k
                // compute it and save it
                memo[start+k] = toPtrBool(dfs(start+k))
                // exit early if we were able to reach
                if *memo[start+k] {return true}
            } else {
                // if we do have an answer saved for start+k
                // return whatever that saved ans is
                return *memo[start+k]
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