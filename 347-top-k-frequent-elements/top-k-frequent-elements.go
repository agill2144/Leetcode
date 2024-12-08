func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    maxFreq := 0
    for i := 0; i < len(nums); i++ {freq[nums[i]]++; maxFreq = max(maxFreq, freq[nums[i]])}
    bucket := make([][]int, maxFreq+1)
    for num , count := range freq {
        if bucket[count] == nil {bucket[count] = []int{}}
        bucket[count] = append(bucket[count], num)
    }
    out := []int{}
    for i := len(bucket)-1; i >= 0 && len(out) != k; i-- {
        out = append(out, bucket[i]...)
    }
    return out
}
