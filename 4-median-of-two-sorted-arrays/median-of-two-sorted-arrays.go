func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {
        return findMedianSortedArrays(nums2, nums1)
    }
    n1 := len(nums1)
    n2 := len(nums2)
    total := n1+n2
    half := total/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        if mid < n1 {r1 = nums1[mid]}
        if mid2 < n2 {r2 = nums2[mid2]}
        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1,l2)) + float64(min(r1,r2))) / 2.0
            }
            return float64(min(r1,r2))
        }

        if l1 > r2 {
            right= mid-1
        } else {
            left = mid+1
        }
    }
    return -1
}
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     totalSize := len(nums1) + len(nums2)
//     n1, n2 := 0, 0
//     // N = totalSize of n1+n2
//     // n = (N/2)+1
//     // tc = o(n)
//     // sc = o(n)
//     for (n1 < len(nums1) || n2 < len(nums2)) && len(merged) < (totalSize/2)+1 {
//         n1Val := math.MaxInt64
//         n2Val := math.MaxInt64
//         if n1 < len(nums1) {n1Val = nums1[n1]}
//         if n2 < len(nums2) {n2Val = nums2[n2]}
//         if n1Val < n2Val {
//             merged = append(merged, nums1[n1])
//             n1++
//         } else {
//             merged = append(merged, nums2[n2])
//             n2++
//         }
//     }
//     if totalSize % 2 == 0 {
//         return (float64(merged[len(merged)-1]) + float64(merged[len(merged)-2])) / 2.0
//     }
//     return float64(merged[len(merged)-1])
// }