func canJump(nums []int) bool {
    farthestIdx := 0
    i := 0
    for i < len(nums) && i <= farthestIdx {
        nextIdx := nums[i] + i
        farthestIdx = max(farthestIdx, nextIdx)
        if farthestIdx >= len(nums)-1 {return true}
        i++
    }
    return false
    
}

// func canJump(nums []int) bool {
//     var dfs func(idx int) bool
//     dfs = func(idx int) bool {
//         // base
//         if idx >= len(nums)-1 {return true}

//         // logic; from current idx, take all possible jumps
//         for k := 1; k <= nums[idx]; k++ {
//             if dfs(idx+k) {return true}
//         }
//         return false
//     }   
//     return dfs(0)
// }