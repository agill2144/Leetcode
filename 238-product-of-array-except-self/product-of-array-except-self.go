func productExceptSelf(nums []int) []int {
    n := len(nums)
    leftProd := make([]int, n)
    leftProd[0] = 1
    for i := 1; i < n; i++ {
        leftProd[i] = nums[i-1]*leftProd[i-1]
    }
    rrp := 1
    for i := n-1; i >= 0; i-- {
        leftProd[i] *= rrp
        rrp *= nums[i]
    }
    return leftProd
}