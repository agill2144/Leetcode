func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1)}
    total := len(nums1) + len(nums2)
    half := (total)/2
    left := 0
    right := len(nums1)
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid < len(nums1) {r1 = nums1[mid]}
        if mid2 < len(nums2) {r2 = nums2[mid2]}
        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1, l2)) + float64(min(r1,r2)))/2.0
            }
            return float64(min(r1, r2))

        }
        if l1 > r2 {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     total := len(nums1) + len(nums2)
//     mid := (total-1)/2
//     mid2 := mid+1
//     undefined := math.MinInt64
//     midVal := undefined
//     mid2Val := undefined
//     idx := 0
//     n1, n2 := 0, 0
//     for n1 < len(nums1) || n2 < len(nums2) {
//         n1Val := math.MaxInt64
//         n2Val := math.MaxInt64
//         if n1 < len(nums1) {n1Val = nums1[n1]}
//         if n2 < len(nums2) {n2Val = nums2[n2]}
//         if n1Val < n2Val {
//             if idx == mid {midVal = n1Val}
//             if idx == mid2 {mid2Val = n1Val; break}
//             n1++
//         } else {
//             if idx == mid {midVal = n2Val}
//             if idx == mid2 {mid2Val = n2Val; break}
//             n2++
//         }
//         idx++
//     }
//     if total % 2 != 0 {
//         return float64(midVal)
//     }
//     return (float64(midVal) + float64(mid2Val)) / 2.0
// }