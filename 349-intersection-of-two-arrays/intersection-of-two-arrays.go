func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
    set := map[int]bool{}
    for i := 0; i < len(nums1); i++ {set[nums1[i]] = true}
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if set[nums2[i]] {
            out = append(out, nums2[i])
            set[nums2[i]] = false
        }
    }
    return out
}