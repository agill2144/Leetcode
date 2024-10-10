func maxWidthRamp(nums []int) int {
    n := len(nums)
    maxRight := make([]int, n)
    maxRight[n-1] = nums[n-1]
    for i := n-2; i >= 0; i-- {
        maxRight[i] = max(maxRight[i+1], nums[i])
    }
    left := 0
    maxWin := 0
    for i := 0; i < n; i++ {
        for nums[left] > maxRight[i] {
            left++
        }
        maxWin = max(maxWin, i-left)

    }
    return maxWin
}