func minSubArrayLen(target int, nums []int) int {
    rSum := 0
    left := 0
    minWin := len(nums)+1
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        for rSum >= target && left <= i {
            minWin = min(minWin, i-left+1)
            rSum -= nums[left]
            left++
        }
    }
    if minWin == len(nums)+1 {return 0}
    return minWin
}