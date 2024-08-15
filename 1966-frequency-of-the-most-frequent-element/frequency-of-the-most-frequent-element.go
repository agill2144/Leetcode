func maxFrequency(nums []int, k int) int {
    sort.Ints(nums)
    n := len(nums)
    right := n-1
    maxWin := 1
    rSum := nums[right]
    for i := n-2; i >= 0; i-- {
        curr := nums[i]
        rSum += curr
        size := right-i+1
        rightVal := nums[right]
        desiredSum := rightVal * size
        diff := desiredSum - rSum
        if diff <= k {
            maxWin = max(maxWin, size)
        } else {
            rSum -= nums[right]
            right--
        }
    }
    return maxWin
}