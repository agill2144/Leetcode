func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {
        return intersect(nums2, nums1)
    }
    
    // nums1 is always smaller
    n1Freq := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        n1Freq[nums1[i]]++
    }
    
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        num := nums2[i]
        val, ok := n1Freq[num]
        if ok && val > 0 {
            out = append(out, num)
            n1Freq[num]--
        }
    }
    return out
}