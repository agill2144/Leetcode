func minSubArrayLen(target int, nums []int) int {
    minWin := len(nums)+1
    sum := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        for sum >= target {
            minWin = min(minWin, i-left+1)
            sum -= nums[left]
            left++
        }
    }
    if minWin == len(nums)+1 {minWin = 0}
    return minWin
}