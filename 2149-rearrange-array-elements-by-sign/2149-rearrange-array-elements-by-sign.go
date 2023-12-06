func rearrangeArray(nums []int) []int {
    positive := 0
    negative := 1
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        if nums[i] >= 0 {
            out[positive] = nums[i]
            positive += 2
        } else {
            out[negative] = nums[i]
            negative += 2
        }
    }
    return out
}