func findKthLargest(nums []int, k int) int {
    n := len(nums)
    start := math.MaxInt64
    end := math.MinInt64
    bucket := map[int]int{}
    for i := 0; i < n; i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
        bucket[nums[i]]++
    }
    idx := -1
    for i := start; i <= end; i++ {
        if bucket[i] == 0 {continue}
        idx += bucket[i]
        if idx >= n-k {return i}
    }
    return -1
}
