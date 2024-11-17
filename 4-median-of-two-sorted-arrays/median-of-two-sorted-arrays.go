func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    merged := []int{}
    merged = append(merged, nums1...)
    merged = append(merged, nums2...)
    sort.Ints(merged)
    n := len(merged)
    mid := (n-1)/2
    if n % 2 != 0 {return float64(merged[mid])}
    return (float64(merged[mid]) + float64(merged[mid+1])) / 2.0
}