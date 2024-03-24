func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    total := len(nums1)+len(nums2)
    mid1 := (total-1)/2
    mid2 := mid1+1
    ele1, ele2 := 0,0
    n1, n2 := 0,0
    idx := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] < nums2[n2] {
            if idx == mid1 {
                ele1 = nums1[n1]
            } else if idx == mid2 {
                ele2 = nums1[n1]
            }
            n1++
            idx++
        } else {
            if idx == mid1 {
                ele1 = nums2[n2]
            } else if idx == mid2 {
                ele2 = nums2[n2]
            }
            n2++
            idx++
        }
    }
    for n1 < len(nums1) {
        if idx == mid1 {
            ele1 = nums1[n1]
        } else if idx == mid2 {
            ele2 = nums1[n1]
        }
        n1++
        idx++
    }
    for n2 < len(nums2) {
        if idx == mid1 {
            ele1 = nums2[n2]
        } else if idx == mid2 {
            ele2 = nums2[n2]
        }
        n2++
        idx++
    }
    if total % 2 == 0 {
        return (float64(ele1) + float64(ele2)) / 2.0
    }
    return float64(ele1)
}

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
