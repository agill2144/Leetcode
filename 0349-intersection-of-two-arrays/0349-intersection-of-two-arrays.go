func intersection(nums1 []int, nums2 []int) []int {
    nums1Set := map[int]bool{}
    for i := 0; i < len(nums1); i++ {
        nums1Set[nums1[i]] = false
    }
    
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        num := nums2[i]
        used, ok := nums1Set[num]
        if ok {
            if !used {
                out = append(out, num)
                nums1Set[num] = true
            }
        }
    }
    return out
}