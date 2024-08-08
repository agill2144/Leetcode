func runningSum(nums []int) []int {
    out := []int{}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        out = append(out, rSum)
    }
    return out
}