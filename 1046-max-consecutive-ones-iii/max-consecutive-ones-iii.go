func longestOnes(nums []int, k int) int {
    zeroCount := 0
    maxSize := 0
    left := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            zeroCount++
        }
        if zeroCount > k {
            if nums[left] == 0 {
                zeroCount--
            }
            left++
            continue
        }
        maxSize = max(maxSize, i-left+1)
    }
    return maxSize
}