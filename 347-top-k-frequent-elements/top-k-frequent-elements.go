func topKFrequent(nums []int, k int) []int {
    n := len(nums)
    freq := map[int]int{}
    for i := 0; i < n; i++ {
        freq[nums[i]]++
    }
    bucket := make([][]int, n+1)
    for k, v := range freq {
        if bucket[v] == nil {bucket[v] = []int{}}
        bucket[v] = append(bucket[v],k)
    }
    out := []int{}
    for i := n; i >= 0 && len(out) != k; i-- {
        if bucket[i] != nil {
            out = append(out, bucket[i]...)
        }
    }
    return out
}

