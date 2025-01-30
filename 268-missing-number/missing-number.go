func missingNumber(nums []int) int {
    n := len(nums)
    ap := (n * (n+1)) / 2
    sum := 0
    for i := 0; i < n; i++ {sum += nums[i]}
    return ap - sum
}