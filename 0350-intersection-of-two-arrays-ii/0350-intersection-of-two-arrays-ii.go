func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    
    n1 := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        n1[nums1[i]]++
    }
    out := []int{}
    for i := 0 ; i < len(nums2); i++ {
        if val, ok := n1[nums2[i]]; ok && val != 0 {
            out = append(out, nums2[i])
            n1[nums2[i]]--
        }
    }
    return out
}