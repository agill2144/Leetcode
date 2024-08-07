func numSubarrayProductLessThanK(nums []int, k int) int {
    if k == 0 {return 0}
    count := 0
    left := 0
    prod := 1
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