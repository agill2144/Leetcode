func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    maxWin := 1
    right := len(nums)-1
    rSum := nums[right]
    for i := len(nums)-2; i >= 0; i-- {
        rSum += nums[i]
        winSize := right-i+1
        desiredSum := winSize * nums[right]
        diff := desiredSum - rSum
        if diff <= k {
            maxWin = max(maxWin, winSize)
        } else {
            rSum -= nums[right]
            right--
        }
    }
    return maxWin
}