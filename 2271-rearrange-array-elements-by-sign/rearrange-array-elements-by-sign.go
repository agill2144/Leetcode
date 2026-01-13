func rearrangeArray(nums []int) []int {
    out := make([]int, len(nums))
    pos := 0
    neg := 1
    for i := 0; i < len(nums); i++ {
        if nums[i] < 0 {
            out[neg] = nums[i]
            neg+=2
        } else {
            out[pos] = nums[i]
            pos+=2
        }
    }
    return out
}