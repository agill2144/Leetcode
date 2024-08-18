/*
    approach: greedy
    - from an ith position, keep track of how far you can jump
    - and then go to next ith position, and update farthest idx as needed
    - keep doing this as long as i <= farthest idx

    time = o(n)
    space = o(1)
*/
func canJump(nums []int) bool {
    farthestIdx := 0
    i := 0
    n := len(nums)
    for i < n && i <= farthestIdx {
        farthestIdx = max(farthestIdx, i+nums[i])
        i++
    }
    return farthestIdx >= n-1
}