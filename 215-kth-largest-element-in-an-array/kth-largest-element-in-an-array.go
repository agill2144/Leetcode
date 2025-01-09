func findKthLargest(nums []int, k int) int {
    start, end := math.MaxInt64, math.MinInt64
    for i := 0; i < len(nums); i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
    }
    // why cant we use a normal bucket array ?
    // because we have negative values
    bucket := map[int]int{}
    for i := 0; i < len(nums); i++ {
        bucket[nums[i]]++
    }
    for i := end; i >= start; i-- {
        if bucket[i] > 0 {k -= bucket[i]; if k <= 0 {return i}}
    }
    return -1
}

/*
3,2,3,1,2,4,5,5,6

1,2,2,3,3,4,5,5,6

1,2,3,4,5,6
*/