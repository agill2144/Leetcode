func rearrangeArray(nums []int) []int {
    pos := 0
    neg := 1
    n := len(nums)
    out := make([]int, n)
    for i := 0; i < n; i++ {
        if nums[i] > 0 {
            out[pos] = nums[i]
            pos+=2
        } else {
            out[neg] = nums[i]
            neg+=2
        }
    }
    return out
}