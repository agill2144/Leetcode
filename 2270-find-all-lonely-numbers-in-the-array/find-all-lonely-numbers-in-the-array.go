func findLonely(nums []int) []int {
    out := []int{}
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {freq[nums[i]]++}
    for key, val := range freq {
        curr := key
        prev := curr-1
        next := curr+1
        if freq[next] == 0 && freq[prev] == 0 && val == 1 {out = append(out, curr)}
    }
    return out
}