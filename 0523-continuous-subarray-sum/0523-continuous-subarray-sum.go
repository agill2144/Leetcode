func checkSubarraySum(nums []int, k int) bool {
    remainderIdx := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        r := sum % k
        idx, ok := remainderIdx[r]
        if ok && i-(idx+1)+1 >= 2 {
            return true
        }
        if !ok {
            remainderIdx[r] = i
        }
    }
    return false
}