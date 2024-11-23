func minOperations(nums []int, x int) int {
    n := len(nums)
    left := n-1
    sum := 0
    for sum < x && left >= 0 {
        sum += nums[left]
        left--
    }
    if sum < x {return -1}
    if left < 0 {left++}
    sum = 0
    i := left
    minWin := math.MaxInt64
    for left <= n {
        sum += nums[i%n]
        for sum > x && left <= n {
            sum -= nums[left%n]
            left++
        }
        if left > n {break}
        if sum == x {
            minWin = min(minWin, i-left+1)
        }
        i++
    }
    if minWin == math.MaxInt64 {minWin = -1}
    return minWin
}