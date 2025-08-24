func maxSubArrayLen(nums []int, k int) int {
    idx := map[int]int{0:-1}
    res := 0
    n := len(nums)
    rSum := 0
    for i := 0; i < n; i++ {
        rSum += nums[i]
        if val, ok := idx[rSum-k]; ok {
            res = max(res, i-val)
        }
        if _, ok := idx[rSum]; !ok {
            idx[rSum] = i
        }
    }
    return res
}