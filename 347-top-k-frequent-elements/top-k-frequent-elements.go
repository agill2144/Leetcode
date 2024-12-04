func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    freqToNums := map[int][]int{}
    maxFreq := 0
    for k, v := range freq {maxFreq = max(maxFreq,v); freqToNums[v] = append(freqToNums[v], k)}
    out := []int{}
    for len(out) != k {
        items := freqToNums[maxFreq]
        for i := 0; i < len(items) && len(out) != k; i++ {out = append(out, items[i])}
        maxFreq--
    }
    return out
}