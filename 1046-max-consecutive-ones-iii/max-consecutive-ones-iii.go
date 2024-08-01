func longestOnes(nums []int, k int) int {
    zeros := 0
    maxWin := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeros++
        }
        for zeros > k {
            if nums[left] == 0 {
                zeros--
            }
            left++
        }
        maxWin = max(maxWin, i-left+1)
    }
    return maxWin
}