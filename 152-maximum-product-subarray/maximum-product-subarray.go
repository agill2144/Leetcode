func maxProduct(nums []int) int {
    n := len(nums)
    leftRp := 1
    rightRp := 1
    res := math.MinInt64
    for i := 0; i < n; i++ {
        leftRp *= nums[i]
        rightRp *= nums[n-1-i]
        res = max(res, max(leftRp, rightRp))
        if leftRp == 0 {leftRp = 1}
        if rightRp == 0 {rightRp = 1}
    }
    return res
}