func getAverages(nums []int, k int) []int {
    n := len(nums)
    out := make([]int,n)
    for i := 0; i < n; i++ {out[i] = -1}
    left := 0
    rSum := 0
    winSize := (k*2)+1
    for i := 0; i < n; i++ {
        rSum += nums[i]
        if i-left+1 == winSize {
            out[i-k] = rSum / winSize
            rSum -= nums[left]
            left++
        }
    }
    return out
}