func maxSubArrayLen(nums []int, k int) int {
    idx := map[int]int{0:-1}
    rSum := 0
    res := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        val, ok := idx[rSum-k]
        if ok {
            res = max(res, i-val)
        }

        if _, ok := idx[rSum]; !ok {
            idx[rSum] = i
        }
    }
    return res
}