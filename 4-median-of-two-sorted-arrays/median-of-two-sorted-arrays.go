func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Val := nums1[n1]
        n2Val := nums2[n2]
        if n1Val <= n2Val {
            merged = append(merged, n1Val)
            n1++
        } else {
            merged = append(merged, n2Val)
            n2++
        }
    }
    for n1 < len(nums1) {
        merged = append(merged, nums1[n1])
        n1++
    }
    for n2 < len(nums2) {
        merged = append(merged, nums2[n2])
        n2++
    }
    mid := (len(merged)-1)/2
    if len(merged) % 2 != 0 {
        return float64(merged[mid])
    }
    return (float64(merged[mid]) + float64(merged[mid+1]))/2.0
}