func sortedSquares(nums []int) []int {
    out := make([]int, len(nums))
    l,r,o := 0, len(nums)-1, len(out)-1
    for o >= 0 {
        lSq := nums[l] * nums[l]
        rSq := nums[r] * nums[r]
        if lSq > rSq {
            out[o] = lSq
            l++
        } else {
            out[o] = rSq
            r--
        }
        o--
    }
    return out
}