func numSubarrayProductLessThanK(nums []int, k int) int {
    count := 0
    prod := 1
    left := 0
    for i := 0; i < len(nums); i++ {
        prod *= nums[i]
        for prod >= k && left <= i {
            prod /= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}