
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1)}
    // nums1 is always smaller in size compared to nums2
    n1 := len(nums1)
    n2 := len(nums2)
    n := n1+n2
    half := n/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        
        l1 := math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        l2 := math.MinInt64
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}        
        r1 := math.MaxInt64
        if mid < n1 {r1 = nums1[mid]}
        r2 := math.MaxInt64
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if n % 2 == 0 {
                return (float64(max(l1,l2)) + float64(min(r1,r2))) / 2.0
            }
            return float64(min(r1,r2))
        }
        if l1 > r2 {
            right = mid-1
        } else{ 
            left = mid+1
        }
    }
    return -1.0
}

// time = o(n1+n2)
// space = o(1)
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     n := len(nums1) + len(nums2) 
//     ele1 := -1
//     idx1 := (n/2)-1
//     ele2 := -1
//     idx2 := n/2
//     idx := 0
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val < n2Val {
//             if idx == idx1 {
//                 ele1 = n1Val
//             } else if idx == idx2 {
//                 ele2 = n1Val
//             }
//             n1++
//         } else {
//             if idx == idx1 {
//                 ele1 = n2Val
//             } else if idx == idx2 {
//                 ele2 = n2Val
//             }
//             n2++
//         }
//         idx++
//     }
//     if ele1 == -1 || ele2 == -1  {
//         for n1 < len(nums1) {
//             n1Val := nums1[n1]
//             if idx == idx1 {
//                 ele1 = n1Val
//             } else if idx == idx2 {
//                 ele2 = n1Val
//             }
//             n1++
//             idx++
//         }

//         for n2 < len(nums2) {
//             n2Val := nums2[n2]
//             if idx == idx1 {
//                 ele1 = n2Val
//             } else if idx == idx2 {
//                 ele2 = n2Val
//             }
//             n2++
//             idx++
//         }
//     }
//     if n % 2 == 0 {
//         return (float64(ele1) + float64(ele2)) / 2.0
//     }
//     return float64(ele2)
// }