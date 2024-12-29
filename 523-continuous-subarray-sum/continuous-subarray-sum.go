func checkSubarraySum(nums []int, k int) bool {
    rems := map[int]int{0:-1}
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        r := sum % k
        if r < 0 {r += k}
        idx, ok := rems[r]
        if ok && i-idx >= 2 {return true}
        if !ok {rems[r] = i}
    }
    return false
}