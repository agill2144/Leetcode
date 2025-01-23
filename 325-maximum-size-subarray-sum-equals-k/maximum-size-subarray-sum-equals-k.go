func maxSubArrayLen(nums []int, k int) int {
    idx := map[int]int{0:-1}
    res := 0
    rSum := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        diff := rSum - k
        prevIdx, ok := idx[diff]
        if ok {
            res = max(res, i-prevIdx)
        }
        if _, ok := idx[rSum]; !ok {
            idx[rSum] = i
        }
    }
    return res
}