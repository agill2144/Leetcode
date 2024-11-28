func longestOnes(nums []int, k int) int {
    zeros := 0
    maxWin := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {zeros++}
        if zeros <= k {
            maxWin = max(maxWin, i-left+1)
        } else {
            if nums[left] == 0 {zeros--}
            left++
        }
    }
    return maxWin
}