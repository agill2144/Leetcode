func maxSubArray(nums []int) int {
    if nums == nil || len(nums) == 0 {return -1}
    if len(nums) == 1 {return nums[0]}
    
    curr := nums[0]
    out := curr

    for i := 1; i < len(nums); i++ {
        curr = max(nums[i], nums[i]+curr)
        out = max(out, curr)
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}