func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {
        return findMedianSortedArrays(nums2, nums1)
    }
    // nums1 is always smaller in size
    total := len(nums1) + len(nums2)
    half := total / 2
    left := 0
    right := len(nums1)
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
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
// noob approach
// we have 2 sorted arrays
// merge them using two ptrs into a third array
// and then give out the median
// n = len(nums1)
// m = len(nums2)
// tc = o(n+m)
// sc = o(n+m)
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val <= n2Val {
//             merged = append(merged, n1Val)
//             n1++
//         } else {
//             merged = append(merged, n2Val)
//             n2++
//         }
//     }
//     for n1 < len(nums1) {
//         merged = append(merged, nums1[n1])
//         n1++
//     }
//     for n2 < len(nums2) {
//         merged = append(merged, nums2[n2])
//         n2++
//     }
//     mid := (len(merged)-1)/2
//     if len(merged) % 2 != 0 {
//         return float64(merged[mid])
//     }
//     return (float64(merged[mid]) + float64(merged[mid+1]))/2.0
// }