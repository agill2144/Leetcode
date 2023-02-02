func findKthLargest(nums []int, k int) int {
    if k > len(nums) {return -1}
    sort.Ints(nums)
    return nums[len(nums)-k]
}