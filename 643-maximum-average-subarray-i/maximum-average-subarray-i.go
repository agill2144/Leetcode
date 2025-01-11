func findMaxAverage(nums []int, k int) float64 {
    res := -math.MaxFloat64
    left := 0
    sum := 0
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if i-left+1 == k {
            res = max(res, float64(sum)/float64(k))
            sum -= nums[left]
            left++
        }
    }
    return res
}