func longestOnes(nums []int, k int) int {
    n := len(nums)
    maxWindow := 0
    left := 0
    zeros := 0
    for i := 0; i < n; i++ {
        if nums[i] == 0 {
            zeros++
        }
        for zeros > k {
            if nums[left] == 0 {
                zeros--
            }
            left++
        }

        maxWindow = max(maxWindow, i-left+1)
    }
    return maxWindow
}