func maxProduct(nums []int) int {
    maxProd := -100000.0
    var rp float64 = 1.0
    var rrp float64 = 1.0
    for i := 0; i < len(nums); i++ {
        rp *= float64(nums[i])
        rrp *= float64(nums[len(nums)-1-i])
        maxProd = max(maxProd, max(rp, rrp))
        if rp == 0.0 {rp = 1.0}
        if rrp == 0.0 {rrp = 1.0}
    }
    return int(maxProd)
}



func max(x, y float64) float64 {
    if x > y {return x}
    return y
}