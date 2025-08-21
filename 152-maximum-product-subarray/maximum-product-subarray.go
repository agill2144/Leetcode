func maxProduct(nums []int) int {
    lrp := 1
    maxP := math.MinInt64
    for i := 0; i < len(nums); i++ {
        lrp *= nums[i]
        maxP = max(maxP, lrp)
        if lrp == 0 {lrp = 1}
    }
    rrp := 1
    for i := len(nums)-1; i >= 0; i-- {
        rrp *= nums[i]
        maxP = max(maxP, rrp)
        if rrp == 0 {rrp = 1}
    }
    return maxP

}