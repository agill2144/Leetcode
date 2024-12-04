func topKFrequent(nums []int, k int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    sort.Slice(nums, func(i, j int)bool{
        if freq[nums[i]] == freq[nums[j]] {
            return nums[i] < nums[j]
        }
        return freq[nums[i]] < freq[nums[j]]
    })
    out := []int{}
    for i := len(nums)-1; i >= 0 && len(out) != k; i-- {
        if len(out) == 0 || out[len(out)-1] != nums[i] {
            out = append(out, nums[i])
        }
    }
    return out
}