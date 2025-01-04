/*
    approach: greedy
    - from an ith position, keep track of how far you can jump
    - and then go to next ith position, and update right idx as needed
    - keep doing this as long as i <= right idx

    time = o(n)
    space = o(1)

    what were we greedy about?
    - as we kept iterating, we kept track of right reachable idx
    - we kept spamming the farthest jump path we can reach (greedily spamming single path to reach an answer)    
*/
func canJump(nums []int) bool {
    n := len(nums)
    if n <= 1 {return true}
    i := 0
    right := 0
    for i <= right {
        newR := i+nums[i]
        i++
        right = max(right, newR)
        if right >= n-1 {return true}
    }
    return false
}
/*
    approach: brute force dfs
    - try all possible jump distances!
        - but lets be a little greedy and take the highest possible jump first
        - if jumpSize is 5
        - instead of doing try jump 1, then 2, then 3... then 5
        - lets start from 5 first
        - we care about whether we can reach end or not
    - start at pos 0
    - then take all possible jumps, and explore each path depth first
    - example : [2,2,0,1]
    - we are idx 0 initially, so take 2 jumps and we end at idx 2 (val = 0)
    - and now at idx 2, explore the same thing
        - but since this is 0, we cannot go anywhere
        - go back to prev recursion
        - do we need to backtrack anything?
        - no, the args are not references, so no
    - and now we are back at idx 0, take the next smaller jump of size 1
    - now we are idx 1, val = 2
    - again, explore all possible jump dists but be greedy and start from highest

    tc:
    - numOfOptionsPerNode^totalNumOfOptions
    - at each node we have x options
    - x is the value of each node in the array
    - tc = x^n
*/
// func canJump(nums []int) bool {
//     n := len(nums)
//     var dfs func(ptr int)bool
//     dfs = func(ptr int)bool {
//         if ptr >= n-1 {return true}
//         jumpSize := nums[ptr]
//         for i := 1; i <= jumpSize; i++ {
//             if dfs(ptr+i) {return true}
//         }
//         return false
//     }
//     return dfs(0)
// }