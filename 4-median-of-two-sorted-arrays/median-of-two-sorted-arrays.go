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
        l1, l2 := math.MinInt64, math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1, r2 := math.MaxInt64, math.MaxInt64
        if mid < n1 { r1 = nums1[mid] }
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (max(float64(l1),float64(l2)) + min(float64(r1),float64(r2))) / 2.0
            }
            return min(float64(r1),float64(r2))
        }

        if l2 > r1 {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return 0.0
}
/*
    approach: optimized space on brute force
    - we just need 2 middle elements in the worst case
        - i.e our merged array could be even or odd
        - worst case its even and we need 2 middle elements
    - so instead of physically placing them into a merged array
    - we can just keep track of 
        1. the current idx we are placing in a fake array
        2. the current element that should be placed at this idx value in fake array
    - if idx is either one of the mid idx we care about (mid or mid+1)
        - we will save its corresponding value
    - then at the end the logic is the same
    - if total size is even
        - avg of both middle elements
    - if total size is off
        - return the first middle element
    
    time = o(n1+n2)
    space = o(1)
    
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     total := len(nums1)+len(nums2)
//     mid1 := (total-1)/2
//     mid2 := mid1+1
//     ele1, ele2 := 0,0
//     n1, n2 := 0,0
//     idx := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             if idx == mid1 {
//                 ele1 = nums1[n1]
//             } else if idx == mid2 {
//                 ele2 = nums1[n1]
//             }
//             n1++
//             idx++
//         } else {
//             if idx == mid1 {
//                 ele1 = nums2[n2]
//             } else if idx == mid2 {
//                 ele2 = nums2[n2]
//             }
//             n2++
//             idx++
//         }
//     }
//     for n1 < len(nums1) {
//         if idx == mid1 {
//             ele1 = nums1[n1]
//         } else if idx == mid2 {
//             ele2 = nums1[n1]
//         }
//         n1++
//         idx++
//     }
//     for n2 < len(nums2) {
//         if idx == mid1 {
//             ele1 = nums2[n2]
//         } else if idx == mid2 {
//             ele2 = nums2[n2]
//         }
//         n2++
//         idx++
//     }
//     if total % 2 == 0 {
//         return (float64(ele1) + float64(ele2)) / 2.0
//     }
//     return float64(ele1)
// }

/*
    approach: brute force
    - using two ptrs, create a 3rd arr
    - find the middle element in the 3rd array
    - if arr size is even
        - median is the avg of 2 middle elements
        - [x,x,x, x,x,x]
    - if arr size is odd
        - median is the middle elment
        - [x,x,x, x, x,x,x]

    n1 = len(nums1)
    n2 = len(nums2)
    n3 = n1+n2
    time = o(n1 + n2)
    space = o(n3)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     sortedArr := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             sortedArr = append(sortedArr, nums1[n1])
//             n1++
//         } else {
//             sortedArr = append(sortedArr, nums2[n2])
//             n2++
//         }
//     }
//     for n1 < len(nums1) {sortedArr = append(sortedArr, nums1[n1]); n1++}
//     for n2 < len(nums2) {sortedArr = append(sortedArr, nums2[n2]); n2++}
//     mid := (len(sortedArr)-1) / 2
//     if len(sortedArr) % 2 != 0 {
//         return float64(sortedArr[mid])
//     }
//     return (float64(sortedArr[mid]) + float64(sortedArr[mid+1])) / 2.0    
// }
