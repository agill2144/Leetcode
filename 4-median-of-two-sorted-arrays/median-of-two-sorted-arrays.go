func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    totalSize := len(nums1) + len(nums2)
    n1, n2 := 0, 0
    // N = totalSize of n1+n2
    // n = (N/2)+1
    // tc = o(n)
    // sc = o(n)
    for (n1 < len(nums1) || n2 < len(nums2)) && len(merged) < (totalSize/2)+1 {
        n1Val := math.MaxInt64
        n2Val := math.MaxInt64
        if n1 < len(nums1) {n1Val = nums1[n1]}
        if n2 < len(nums2) {n2Val = nums2[n2]}
        if n1Val < n2Val {
            merged = append(merged, nums1[n1])
            n1++
        } else {
            merged = append(merged, nums2[n2])
            n2++
        }
    }
    if totalSize % 2 == 0 {
        return (float64(merged[len(merged)-1]) + float64(merged[len(merged)-2])) / 2.0
    }
    return float64(merged[len(merged)-1])
}