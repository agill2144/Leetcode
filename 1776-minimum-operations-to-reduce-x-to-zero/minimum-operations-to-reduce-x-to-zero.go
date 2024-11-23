func minOperations(nums []int, x int) int {
    n := len(nums)
    winSum := nums[n-1]
    i := n-1
    for i >= 0 && winSum < x {
        winSum += nums[i]
        i--
    }
    if i == -1 {i = 0}
    if winSum < x {return -1}
    winSum = 0
    left := i
    minWin := len(nums)+1
    for left <= n {
        winSum += nums[i%n]
        for winSum > x && left <= n {
            winSum -= nums[left%n]
            left++
        }
        if left > n {break}
        if winSum == x { minWin = min(minWin, i-left+1) }
        i++
    }
    if minWin == len(nums)+1 {return -1}
    return minWin
}