func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) > len(nums1) {
        return intersect(nums2, nums1)
    }
    
    // nums1 is always the larger array
    freqMap := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        freqMap[nums1[i]]++
    }
    
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if val, ok := freqMap[nums2[i]]; ok && val != 0 {
            out = append(out, nums2[i])
            freqMap[nums2[i]]--
        }
    }
    return out
}