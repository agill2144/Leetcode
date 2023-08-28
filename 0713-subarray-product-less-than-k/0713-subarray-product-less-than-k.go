func numSubarrayProductLessThanK(nums []int, k int) int {
    if nums == nil || len(nums) == 0 || k <= 1 {
        return 0
    }

    left := 0
    rp := 1
    count := 0
    for i := 0; i < len(nums); i++ {
        rp *= nums[i]
        for rp >= k {
            rp = rp / nums[left]
            left++
        }
        count += i-left+1
    }
    return count
}