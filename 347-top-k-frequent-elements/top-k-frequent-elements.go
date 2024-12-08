func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    sort.Slice(nums, func(i, j int)bool{
        iFreq := freq[nums[i]]
        jFreq := freq[nums[j]]
        if iFreq == jFreq {
            return nums[i] < nums[j]
        }
        return iFreq < jFreq
    })
    out := []int{}
    for i := len(nums)-1; i >= 0 && len(out) != k; i-- {
        if len(out) != 0 && out[len(out)-1] == nums[i] {continue}
        out = append(out, nums[i])
    }
    return out
}
