func intersect(nums1 []int, nums2 []int) []int {
    n1Freq := map[int]int{}
    for i := 0; i < len(nums1); i++ {
        n1Freq[nums1[i]]++
    }
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        if val := n1Freq[nums2[i]]; val > 0 {
            out = append(out, nums2[i])
            n1Freq[nums2[i]]--
        }
    }
    return out
}