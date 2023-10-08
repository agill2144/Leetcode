func maxProduct(nums []int) int {
    prefix := 1
    suffix := 1
    out := math.MinInt64
    for i := 0; i < len(nums); i++  {
        if prefix == 0 {prefix=1}
        if suffix == 0 {suffix=1}

        prefix *= nums[i]            
        suffix *= nums[len(nums)-1-i]
        out = max(out, max(prefix, suffix))
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}