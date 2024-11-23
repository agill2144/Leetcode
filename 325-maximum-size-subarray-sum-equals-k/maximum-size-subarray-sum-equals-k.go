func maxSubArrayLen(nums []int, k int) int {
    idx := map[int]int{0:-1}
    rSum := 0
    maxSize := 0
    for i := 0; i < len(nums); i++ {
        rSum += nums[i]
        left , ok := idx[rSum-k]
        if ok {
            maxSize = max(maxSize, i-left)
        }
        if _, ok := idx[rSum]; !ok {idx[rSum] = i}
    }
    return maxSize
}