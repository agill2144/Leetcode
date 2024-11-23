func checkSubarraySum(nums []int, k int) bool {
    if len(nums) < 2 {return false}
    rem := map[int]int{0:-1}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        r := rSum % k
        left, ok := rem[r]
        if ok {
            // left++ // from left+1 idx, we have a sum who is divisible by k, therefore left++
            if i-left >= 2 {return true}
        } else {
            rem[r] = i
        }
    }
    return false
}