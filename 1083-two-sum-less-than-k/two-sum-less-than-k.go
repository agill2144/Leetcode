func twoSumLessThanK(nums []int, k int) int {
    sort.Ints(nums)
    left := 0
    right := len(nums)-1
    maxSum := -1
    for left < right {
        sum := nums[left]+nums[right]
        if sum < k {
            maxSum = max(maxSum, sum)
            left++
        } else {
            right--
        }
    }
    return maxSum
}