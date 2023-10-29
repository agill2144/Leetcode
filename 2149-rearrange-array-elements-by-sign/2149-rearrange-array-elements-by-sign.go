func rearrangeArray(nums []int) []int {
    pos := 0
    neg := 1
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if nums[i] > 0 {
            out[pos] = nums[i]
            pos+=2
        } else {
            out[neg] = nums[i]
            neg += 2
        }
    }
    return out
}