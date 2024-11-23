func numSubarrayProductLessThanK(nums []int, k int) int {
    if k == 0 {return 0}
    left := 0
    prod := 1
    count := 0
    for i := 0; i < len(nums); i++ {
        prod *= nums[i]
        for left <= i && prod >= k {
            prod /= nums[left]
            left++
        }
        count += (i-left+1)
    }
    return count
}