func merge(nums1 []int, m int, nums2 []int, n int)  {
    merged := []int{}
    n1 := 0
    n2 := 0
    for n1 < m && n2 < len(nums2) {
        if nums1[n1] < nums2[n2] {
            merged = append(merged, nums1[n1])
            n1++
        } else {
            merged = append(merged, nums2[n2])
            n2++
        }
    }
    for n1 < m {merged = append(merged, nums1[n1]); n1++}
    for n2 < len(nums2) {merged = append(merged, nums2[n2]); n2++}
    for i := 0; i < len(merged); i++ {
        nums1[i] = merged[i]
    }
    
}