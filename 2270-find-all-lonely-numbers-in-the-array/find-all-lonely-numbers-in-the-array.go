func findLonely(nums []int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    out := []int{}
    for k, v := range freq {
        if v > 1 {continue}
        if freq[k-1] == 0 && freq[k+1] == 0 {
            out = append(out, k)
        }
    }
    return out
}