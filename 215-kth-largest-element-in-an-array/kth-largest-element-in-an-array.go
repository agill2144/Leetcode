func findKthLargest(nums []int, k int) int {
    n := len(nums)
    sort.Ints(nums)
    return nums[n-k]
}