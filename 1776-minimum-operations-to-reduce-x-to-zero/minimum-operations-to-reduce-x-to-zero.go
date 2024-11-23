func minOperations(nums []int, x int) int {
    n := len(nums)
    sum := nums[n-1]
    i := n-1
    for i >= 0 && sum < x {
        sum += nums[i]
        i--
    }
    if i == -1 {i = 0}
    if sum < x {return -1}
    sum = 0
    left := i
    minWin := len(nums)+1
    for left <= n {
        sum += nums[i%n]
        for sum > x && left <= n {
            sum -= nums[left%n]
            left++
        }
        if left > n {break}
        if sum == x { minWin = min(minWin, i-left+1) }
        i++
    }
    if minWin == len(nums)+1 {return -1}
    return minWin
}