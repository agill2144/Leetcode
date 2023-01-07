func maxSubArray(nums []int) int {
    curr := nums[0]
    out := nums[0]    
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