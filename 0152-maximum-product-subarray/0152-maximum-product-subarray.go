func maxProduct(nums []int) int {
    prefix := 1
    suffix := 1
    out := math.MinInt64
    
    for i := 0; i < len(nums); i++ {
        prefix *= nums[i]
        suffix *= nums[len(nums)-1-i]
        out = max(out, max(prefix, suffix))
        if prefix == 0{
            // we multiplied with 0
            // reset and start forming a new subarray
            prefix = 1
        }
        if suffix == 0 {
            // we multiplied with 0
            // reset and start forming a new subarray
            suffix = 1
        }
    }
    return out
}

func max(x, y int) int {
    if x > y {return x}
    return y
}