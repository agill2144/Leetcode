func maxProduct(nums []int) int {
    rp := 1
    brp := 1
    out := math.MinInt64
    for i := 0; i < len(nums); i++ {
        rp *= nums[i]
        brp *= nums[len(nums)-1-i]
        out = max(out, max(rp, brp))
        if rp == 0 {rp = 1}
        if brp == 0 {brp = 1}
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}