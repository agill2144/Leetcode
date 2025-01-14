func productExceptSelf(nums []int) []int {
    n := len(nums)
    leftProd := make([]int, n)
    leftProd[0] = 1
    for i := 1; i < n; i++ {
        leftProd[i] = nums[i-1]*leftProd[i-1]
    }
    rightProd := make([]int, n)
    rightProd[n-1] = 1
    for i := n-2; i >= 0; i-- {
        rightProd[i] = rightProd[i+1] * nums[i+1]
    }
    out := make([]int, n)
    for i := 0; i < n; i++{
        out[i] = leftProd[i] * rightProd[i]
    }
    return out
}