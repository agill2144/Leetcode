/*
    approach: greedy
    - from an ith position, keep track of how far you can jump
    - and then go to next ith position, and update farthest idx as needed
    - keep doing this as long as i <= farthest idx

    time = o(n)
    space = o(1)

    what were we greedy about?
    - as we kept iterating, we kept track of farthest reachable idx
    - we kept spamming the farthest jump path we can reach (greedily spamming single path to reach an answer)    
*/
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