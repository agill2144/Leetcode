func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    maxWin := 1
    right := n-1
    rSum := nums[n-1]
    for i := n-2; i >= 0; i-- {
        size := right-i+1
        val := nums[right]
        desiredSum := val * size
        rSum += nums[i]
        if desiredSum - rSum <= k {
            maxWin = max(maxWin, size)
            continue
        } else {
            rSum -= nums[right]
            right--
        }
    }
    return maxWin
}