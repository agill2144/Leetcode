func intersect(nums1 []int, nums2 []int) []int {
    freq := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        freq[nums1[i]]++
    }
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if freq[nums2[i]] > 0 {
            out = append(out, nums2[i])
            freq[nums2[i]]--
        }
    }
    return out
}