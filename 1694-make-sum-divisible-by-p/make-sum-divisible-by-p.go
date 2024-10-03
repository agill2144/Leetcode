func minSubarray(nums []int, p int) int {
    total := 0
    for i := 0; i < len(nums); i++ {
        total += nums[i]
    }
    k := total % p
    if k == 0 {return 0}
    minWin := len(nums)
    rSum := 0
    idx := map[int]int{0:-1}
    for i := 0; i < len(nums); i++ {
        rSum = (rSum + nums[i]) % p
        diff := (rSum - k + p) % p 
        if left, ok := idx[diff]; ok {
            if i-left < minWin {minWin = i-left}
        }
        idx[rSum] = i
    }
    if minWin == len(nums) {minWin = -1}
    return minWin
}