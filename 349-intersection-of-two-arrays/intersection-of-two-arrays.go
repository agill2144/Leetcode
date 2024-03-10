// hash hash
func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
    // nums1 is always smaller
    n1Set := map[int]struct{}{}
    for i := 0; i < len(nums1); i++ {
        n1Set[nums1[i]] = struct{}{}
    }
    outSet := map[int]struct{}{}
    for i := 0; i < len(nums2); i++ {
        num := nums2[i]
        if _, ok := outSet[num]; ok {continue}
        if _, ok := n1Set[num]; ok {
            outSet[num] = struct{}{}
        }
    }
    out := []int{}
    for k, _ := range outSet {
        out = append(out, k)
    }
    return out
}