func longestOnes(nums []int, k int) int {
    left := 0
    zeros := 0
    maxWindow := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeros++
        }
        if zeros > k {
            if nums[left] == 0 {
                zeros--
            }
            left++
        }
        maxWindow = max(maxWindow, i-left+1)
    }
    return maxWindow
}