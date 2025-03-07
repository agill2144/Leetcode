func intersection(nums1 []int, nums2 []int) []int {
    freq := map[int]bool{}
    for i := 0; i < len(nums1); i++ {
        freq[nums1[i]] = true
    }
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if freq[nums2[i]] {
            delete(freq,nums2[i])
            out = append(out, nums2[i])
        }
    }
    return out
}