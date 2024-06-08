func checkSubarraySum(nums []int, k int) bool {
    // find subarray, whose sum is divisble by k
    // subarray size must be >= 2
    remainders := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        idx, ok := remainders[rem]
        if ok {
            if i-idx >= 2 {return true}
        } else {
            remainders[rem] = i
        }
    }
    return false
}