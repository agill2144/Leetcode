func maxProduct(nums []int) int {
    prefix := 1
    suffix := 1
    maxProd := math.MinInt64
    for i := 0; i < len(nums); i++ {
        prefix *= nums[i]
        suffix *= nums[len(nums)-1-i]
        maxProd = max(maxProd, max(prefix, suffix))
        if prefix == 0 {prefix = 1}
        if suffix == 0 {suffix = 1}
    }
    return maxProd
}

func max(x, y int) int {
    if x > y {return x}
    return y
}