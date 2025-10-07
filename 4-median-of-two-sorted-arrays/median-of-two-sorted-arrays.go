func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1) + len(nums2)
    mid := (total-1)/2
    mid2 := mid+1
    undefined := math.MinInt64
    midVal := undefined
    mid2Val := undefined
    idx := 0
    n1, n2 := 0, 0
    for n1 < len(nums1) || n2 < len(nums2) {
        n1Val := math.MaxInt64
        n2Val := math.MaxInt64
        if n1 < len(nums1) {n1Val = nums1[n1]}
        if n2 < len(nums2) {n2Val = nums2[n2]}
        if n1Val < n2Val {
            if idx == mid {midVal = n1Val}
            if idx == mid2 {mid2Val = n1Val; break}
            n1++
        } else {
            if idx == mid {midVal = n2Val}
            if idx == mid2 {mid2Val = n2Val; break}
            n2++
        }
        idx++
    }
    if total % 2 != 0 {
        return float64(midVal)
    }
    return (float64(midVal) + float64(mid2Val)) / 2.0
}