func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return 0}
    curr := nums[0]
    out := nums[0]

    for i := 1; i < len(nums); i++ {
        curr = max(curr+nums[i], nums[i])
        out = max(out, curr)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}