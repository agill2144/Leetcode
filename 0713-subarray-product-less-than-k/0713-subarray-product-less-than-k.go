func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {return 0}
    prod := 1
    left := 0
    count := 0
    for right := 0; right < len(nums); right++ {
        prod *= nums[right]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        count += right-left+1
    }
    return count
}