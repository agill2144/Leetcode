func minOperations(nums []int, x int) int {
    n := len(nums)
    k := 0
    for i := 0; i < n; i++ {k += nums[i]}
    k -= x
    left := 0
    rSum := 0
    minWin := len(nums)+1
    for i := 0; i < n; i++ {
        rSum += nums[i]
        for rSum > k && left <= i {
            rSum -= nums[left]
            left++
        }
        if rSum == k {
            minWin = min(minWin, n-(i-left+1))            
        }
    }
    if minWin == len(nums)+1 {minWin = -1}
    return minWin 
}