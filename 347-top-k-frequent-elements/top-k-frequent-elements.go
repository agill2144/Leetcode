func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    sort.Slice(nums, func(i, j int)bool{
        if freq[nums[i]] == freq[nums[j]] {
            return nums[i] < nums[j]
        }
        return freq[nums[i]] < freq[nums[j]]
    })
    out := []int{}
    for i := len(nums)-1; i >= 0 && len(out) != k; i-- {
        lastNum := math.MinInt64
        if len(out) != 0 {lastNum = out[len(out)-1]}
        if nums[i] != lastNum {
            out = append(out, nums[i])
        }
    }
    return out
}