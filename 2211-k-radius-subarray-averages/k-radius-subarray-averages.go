func getAverages(nums []int, k int) []int {
    n := len(nums)
    kk := k*2+1
    out := make([]int, n)
    left := 0
    sum := 0
    for i := 0; i < n; i++ {
        out[i] = -1
        sum += nums[i]
        if i-left+1 == kk {
            idx := i-k
            avg := sum / kk
            out[idx] = avg
            sum -= nums[left]
            left++
        }
    }
    return out
}