func smallerNumbersThanCurrent(nums []int) []int {
    start := math.MaxInt64
    end := math.MinInt64
    for i := 0; i < len(nums); i++ {
        start = min(start, nums[i])
        end = max(end, nums[i])
    }
    bucket := make([]int, end+1)
    for i := 0; i < len(nums); i++ {
        bucket[nums[i]]++
    }
    prefix := make([]int, end+1)
    for i := start; i < len(bucket); i++ {
        prevCount := 0
        prevPrefixCount := 0
        if i-1 >= 0 {prevCount = bucket[i-1]; prevPrefixCount=prefix[i-1]}
        prefix[i]= prevCount+prevPrefixCount

    }
    out := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        out[i] = prefix[nums[i]]
    }
    return out
}