func numSubarrayProductLessThanK(nums []int, k int) int {
    if k <= 1 {return 0}
    prod := 1
    left := 0
    count := 0
    for i := 0; i < len(nums); i++ {
        prod *= nums[i]
        for prod >= k {
            prod /= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}