func checkSubarraySum(nums []int, k int) bool {
    rems := map[int]int{0:-1}
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        r := rSum % k
        if r < 0 { r += k}
        left, ok := rems[r]
        if ok {
            if i-left >= 2 {return true}
        } else {
            rems[r] = i
        }
    }
    return false
}