func checkSubarraySum(nums []int, k int) bool {
    rIdx := map[int]int{0:-1}
    sum := 0
    
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        idx, ok := rIdx[rem]
        if ok {
            if i-idx >= 2 {
                return true
            }
        } else {
            rIdx[rem] = i
        }
    }
    return false
}