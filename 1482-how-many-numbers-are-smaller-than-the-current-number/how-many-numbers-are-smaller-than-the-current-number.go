func smallerNumbersThanCurrent(nums []int) []int {
    start := math.MaxInt64
    end := math.MinInt64
    n := len(nums)
    for i := 0; i < n ; i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
    }
    
    bucket := make([]int, end+1)
    for i := 0; i < n ; i++ {
        bucket[nums[i]]++
    }
    
    countMap := map[int]int{}
    prev := -1
    out := make([]int, n)
    for i := 0; i < len(bucket); i++ {
        if bucket[i] == 0 { continue }
        prevCountInCountMap := 0
        prevCountInBucket := 0
        if prev != -1 {
            prevCountInCountMap = countMap[prev]
            prevCountInBucket = bucket[prev]
        }
        countMap[i] = prevCountInBucket + prevCountInCountMap
        prev = i
    }
    
    for i := 0; i < n; i++ {
        out[i] = countMap[nums[i]]
    }
    return out
}