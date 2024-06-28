// func maxProduct(nums []int) int {
//     maxProd := math.MinInt64
//     rp := 1
//     rrp := 1
//     for i := 0; i < len(nums); i++ {
//         rp *= nums[i]
//         rrp *= nums[len(nums)-1-i]
//         maxProd = max(maxProd, max(rp, rrp))
//         if rp == 0 {rp = 1}
//         if rrp == 0 {rrp = 1}
//     }
//     return maxProd
// }


func maxProduct(nums []int) int {
    var prefix float64 = 1.0
    var suffix float64 = 1.0
    var out float64 = -100000.0
    
    for i := 0; i < len(nums); i++ {
        prefix *= float64(nums[i])
        suffix *= float64(nums[len(nums)-1-i])
        out = max(out, max(prefix, suffix))
        if prefix == 0.0{
            // we multiplied with 0
            // reset and start forming a new subarray
            prefix = 1.0
        }
        if suffix == 0.0 {
            // we multiplied with 0
            // reset and start forming a new subarray
            suffix = 1.0
        }
    }
    return int(out)
}

func max(x, y float64) float64 {
    if x > y {return x}
    return y
}