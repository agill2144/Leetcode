func longestOnes(nums []int, k int) int {
    maxWin := 0
    left := 0
    zeroCount := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {
            zeroCount++
        }
        for zeroCount > k {
            if nums[left] == 0 {zeroCount--}
            left++
        }
        if zeroCount <= k {
            maxWin = max(maxWin, i-left+1)
        }
    }
    return maxWin
}