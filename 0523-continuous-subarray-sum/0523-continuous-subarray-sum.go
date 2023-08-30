func checkSubarraySum(nums []int, k int) bool {
    remainderIdx := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        remainder := sum % k
        if remainder == 0 && i >= 1 {return true}
        val, ok := remainderIdx[remainder]
        if ok && i-(val+1)+1 >= 2 {
            return true
        }
        if !ok {
            remainderIdx[remainder] = i
        }
    }
    return false
}