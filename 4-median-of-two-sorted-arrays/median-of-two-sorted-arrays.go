/*
    approach: sorted = binary search
    let n1 be the smaller array

    time = o(log min(n1,n2))
    space = o(1)
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n2 < n1 {
        return findMedianSortedArrays(nums2, nums1)
    }
    total := n1+n2
    half := total/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid

        l1, l2 := math.MinInt64,math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1, r2 := math.MaxInt64,math.MaxInt64
        if mid < n1 {r1 = nums1[mid]}
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1,l2)) + float64(min(r1,r2))) / 2.0
            }
            return float64(min(r1,r2))
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
//     n1, n2 := len(nums1), len(nums2)
//     n1Ptr, n2Ptr := 0,0
//     total := n1+n2
//     mid1 := (total-1)/2
//     ele1 := math.MinInt64
//     mid2 := mid1+1
//     ele2 := math.MinInt64
//     idx := 0
//     for n1Ptr < n1 && n2Ptr < n2 {
//         n1Val := nums1[n1Ptr]
//         n2Val := nums2[n2Ptr]
//         if idx == mid1 {
//             ele1 = min(n1Val, n2Val)
//         } else if idx == mid2 {
//             ele2 = min(n1Val, n2Val)
//         }
//         if n1Val < n2Val {
//             n1Ptr++
//         } else {
//             n2Ptr++
//         }
//         idx++
//     }
//     for n1Ptr < n1 {
//         n1Val := nums1[n1Ptr]
//         if idx == mid1 {
//             ele1 = n1Val
//         } else if idx == mid2 {
//             ele2 = n1Val
//         }
//         n1Ptr++
//         idx++
//     }

//     for n2Ptr < n2 {
//         n2Val := nums2[n2Ptr]
//         if idx == mid1 {
//             ele1 = n2Val
//         } else if idx == mid2 {
//             ele2 = n2Val
//         }
//         n2Ptr++
//         idx++
//     }

//     if total % 2 == 0 {
//         return (float64(ele1) + float64(ele2)) / 2.0
//     }
//     return float64(ele1)
// }
/*
    approach: brute force
    - merge 2 sorted arrays into a third array
    - then median is the middle element in the merged array
    - if the merged array size is even
        - median lies in between 2 elements
        - i.e its a decimal number
        - i.e mid + mid+1 / 2
        - i.e avg of 2 middle elements
    - if the merged array size is odd
        - median is the middle elment
        - return mid value
    
    time = o(n1 + n2)
    space = o(n1+n2)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     n1, n2 := len(nums1), len(nums2)
//     n1Ptr, n2Ptr := 0,0
    
//     for n1Ptr < n1 && n2Ptr < n2 {
//         n1Val := nums1[n1Ptr]
//         n2Val := nums2[n2Ptr]
//         if n1Val < n2Val {
//             merged = append(merged, n1Val)
//             n1Ptr++
//         } else {
//             merged = append(merged, n2Val)
//             n2Ptr++
//         }
//     }

//     for n1Ptr < n1 {
//         n1Val := nums1[n1Ptr]
//         merged = append(merged, n1Val)
//         n1Ptr++
//     }
//     for n2Ptr < n2 {
//         n2Val := nums2[n2Ptr]
//         merged = append(merged, n2Val)
//         n2Ptr++
//     }
    
//     mid := (len(merged)-1)/2
//     if len(merged) % 2 == 0 {
//         // median lies in between 2 middle elements
//         mid1 := float64(merged[mid])
//         mid2 := float64(merged[mid+1])
//         return (mid1+mid2) / 2.0
//     }
//     return float64(merged[mid])
// }