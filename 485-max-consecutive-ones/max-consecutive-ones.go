func findMaxConsecutiveOnes(nums []int) int {
    ones := 0
    left := 0
    maxWin := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            if left == -1 {left = i}
            ones++
            size := i-left+1
            maxWin = max(maxWin, size)
        } else {
            left = -1
            ones = 0
        }
    }
    return maxWin
}