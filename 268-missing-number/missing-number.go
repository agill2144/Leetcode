
func missingNumber(nums []int) int {
    sum := 0
    for i := 0; i < len(nums); i++ {sum += nums[i]}
    n := len(nums)
    expected := n*(n+1)/2
    return expected-sum
}