func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n2 < n1 {return findMedianSortedArrays(nums2, nums1)}
    total := n1+n2
    half := total/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid < n1 {r1 = nums1[mid]}
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2.0
            }
            return float64(min(r1, r2))
        }
        if l2 > r1 {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return 0.0
}

// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     n := len(nums1) + len(nums2)
//     mid1 := (n-1)/2
//     ele1 := -1
//     mid2 := mid1+1
//     ele2 := -1
//     idx := 0
//     n1 := 0; n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val < n2Val {
//             if idx == mid1 {ele1 = n1Val}
//             if idx == mid2 {ele2 = n1Val}
//             idx++
//             n1++
//         } else {
//             if idx == mid1 {ele1 = n2Val}
//             if idx == mid2 {ele2 = n2Val}
//             idx++
//             n2++
//         }
//     }
//     for n1 < len(nums1){
//         n1Val := nums1[n1]
//         if idx == mid1 {ele1 = n1Val}
//         if idx == mid2 {ele2 = n1Val}
//         idx++
//         n1++
//     }

//     for n2 < len(nums2){
//         n2Val := nums2[n2]
//         if idx == mid1 {ele1 = n2Val}
//         if idx == mid2 {ele2 = n2Val}
//         idx++
//         n2++
//     }

//     if n % 2 != 0 {return float64(ele1)}
//     return (float64(ele1) + float64(ele2)) / 2.0

// }