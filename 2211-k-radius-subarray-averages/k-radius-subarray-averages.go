func getAverages(nums []int, k int) []int {
    out := make([]int, len(nums))
    kk := k*2+1
    left := 0
    rSum := 0
    for i := 0; i < len(nums); i++ {
        out[i] = -1
        rSum += nums[i]
        if i-left+1 == kk {
            out[i-k] = rSum / kk
            rSum -= nums[left]
            left++
        }
    }
    return out
}