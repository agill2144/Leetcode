func checkSubarraySum(nums []int, k int) bool {
    remainderIdx := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        rem := sum % k
        if rem < 0 {rem += k}
        idx, ok := remainderIdx[rem]
        if ok && i-(idx+1)+1 >= 2 {return true}
        if !ok {
            remainderIdx[rem] = i
        }
    }
    return false
}