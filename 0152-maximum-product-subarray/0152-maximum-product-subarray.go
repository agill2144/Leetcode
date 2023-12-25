func maxProduct(nums []int) int {
    rp := 1
    rrp := 1 // running prod from right side of array
    out := math.MinInt64
    for i := 0; i < len(nums); i++ {
        rp *= nums[i]
        rrp *= nums[len(nums)-1-i]
        out = max(out, max(rp, rrp))
        if rp == 0 {rp = 1}
        if rrp == 0 {rrp = 1}
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}