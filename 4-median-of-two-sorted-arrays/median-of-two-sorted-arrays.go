func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    totalSize := len(nums1) + len(nums2)
    n1, n2 := 0, 0
    for n1 < len(nums1) && n2 < len(nums2) && len(merged) < (totalSize/2)+1 {
        if nums1[n1] < nums2[n2] {
            merged = append(merged, nums1[n1])
            n1++
        } else {
            merged = append(merged, nums2[n2])
            n2++
        }
        if len(merged) > (totalSize/2)+1 {break}
    }
    for n1 < len(nums1) && len(merged) < (totalSize/2)+1 {
        merged = append(merged, nums1[n1])
        n1++
    }
    for n2 < len(nums2) && len(merged) < (totalSize/2)+1 {
        merged = append(merged, nums2[n2])
        n2++
    }
    if totalSize % 2 == 0 {
        return (float64(merged[len(merged)-1]) + float64(merged[len(merged)-2])) / 2.0
    }
    return float64(merged[len(merged)-1])
}