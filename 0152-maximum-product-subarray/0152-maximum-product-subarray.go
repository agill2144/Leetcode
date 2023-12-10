func maxProduct(nums []int) int {
    max := math.MinInt64
    rp := 1
    for i := 0; i < len(nums); i++ {
        for j := i; j < len(nums); j++ {
            rp *= nums[j]
            if rp > max {max = rp}
        }
        rp = 1
    }
    return max
}