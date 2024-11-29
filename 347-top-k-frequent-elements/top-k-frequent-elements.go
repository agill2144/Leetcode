func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}

    bucket := make([][]int, len(nums)+1)
    for num, count := range freq {
        if bucket[count] == nil {bucket[count] = []int{}}
        bucket[count] = append(bucket[count], num)
    }
    out := []int{}
    for i := len(bucket)-1; i >= 0 && len(out) != k; i-- {
        if bucket[i] == nil {continue}
        out = append(out, bucket[i]...)
    }
    return out
}
