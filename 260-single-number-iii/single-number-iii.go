func singleNumber(nums []int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
    }
    out := []int{}
    for k,v := range freq {
        if v == 1 {out = append(out, k)}
    }
    return out
}