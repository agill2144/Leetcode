func buildArray(nums []int) []int {
    out := make([]int, len(nums)) 
    for i := 0; i < len(nums); i++ {
        out[i] = nums[nums[i]]
    }
    return out
}