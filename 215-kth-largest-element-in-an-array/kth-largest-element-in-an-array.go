// bucket sort / counting sort
// use map as a bucket instead of an array
// why ? we have negative values, and we cannot easily plot/map a negative value to a idx in an array
// therefore using map
// time = o(n) + o(n) + o(max-min+1) 
// space = o(n)
func findKthLargest(nums []int, k int) int {
    n := len(nums)
    if k > n {return -1}

    start := math.MaxInt64
    end := math.MinInt64
    for i := 0; i < n; i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
    }

    bucket := map[int]int{}
    for i := 0; i < n; i++ {
        bucket[nums[i]]++
    }

    idx := -1
    targetIdx := n-k
    for i := start; i <= end; i++ {
        count, ok := bucket[i]
        if !ok {continue}
        idx += count
        if idx >= targetIdx {return i}
    }
    return -1
}