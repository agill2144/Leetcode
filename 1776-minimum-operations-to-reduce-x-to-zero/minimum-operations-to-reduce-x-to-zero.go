func minOperations(nums []int, x int) int {
    n := len(nums)
    if nums[0] > x && nums[n-1] > x {return -1} 
    total := 0
    for i := 0; i < n; i++ {total += nums[i]}
    k := total-x
    winSum := 0
    minWin := len(nums)+1
    left := 0
    for i := 0; i < n; i++ {
        winSum += nums[i]
        for left <= i && winSum > k {
            winSum -= nums[left]
            left++
        }
        if winSum == k {
            minWin = min(minWin, n-(i-left+1))
        }
    }
    if minWin == len(nums)+1 {return -1}
    return minWin
}