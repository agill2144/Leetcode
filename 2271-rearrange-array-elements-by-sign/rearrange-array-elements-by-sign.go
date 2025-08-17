func rearrangeArray(nums []int) []int {
    n := len(nums)
    out := make([]int, n)
    pos := 0
    neg := 1
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            out[pos] = nums[i]
            pos += 2
        } else {
            out[neg] = nums[i]
            neg += 2
        }
    }
    return out

}