func anagramMappings(nums1 []int, nums2 []int) []int {
    nums2Idx := map[int]int{}
    for i := 0; i < len(nums2); i++ {
        nums2Idx[nums2[i]] = i
    }
    out := make([]int, len(nums1))
    for i := 0; i < len(nums2); i++ {
        out[i] = nums2Idx[nums1[i]]
    }
    return out
    
}