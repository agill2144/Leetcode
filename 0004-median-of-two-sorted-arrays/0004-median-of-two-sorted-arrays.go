/*
    approach: binary search
    time: o(log nums1)
    space: o(1)
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1)}
    
    n1 := len(nums1)
    n2 := len(nums2)
    
    left := 0
    right := n1
    
    /*
        [1,4,6,8]
        [2,5,7,9,10,11,12,13]
    */
    for left <= right {
        partX := left + (right-left) / 2
        partY := (n1+n2)/2 - partX
        
        l1 := math.MinInt64
        if partX != 0 { l1 = nums1[partX-1]}
        l2 := math.MinInt64
        if partY != 0 { l2 = nums2[partY-1]}
        
        r1 := math.MaxInt64
        if partX < n1 { r1 = nums1[partX]}
        r2 := math.MaxInt64
        if partY < n2 { r2 = nums2[partY]}
        
        if l1 <= r2 && l2 <= r1 {
            if (n1+n2) % 2 == 0 {
                return (math.Min(float64(r1), float64(r2)) + math.Max(float64(l1), float64(l2)))/2.0
            } else {
                return float64(min(r1,r2))
            }
        } else if l1 > r2 {
            right = partX-1
        } else {
            left = partX+1
        }
    }
    return -1
}

func min(x, y int) int {
    if x < y {return x}
    return y
}

func max(x, y int) int {
    if x > y {return x}
    return y
}



/*
    approach : merge + sort and then grab middle
    - merge both arrays into a new m+n array
    - Sort the new array
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+nlogm+n)
    space: o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     for i := 0; i < len(nums1); i++ {
//         merged = append(merged, nums1[i])
//     }
//     for i := 0; i < len(nums2); i++ {
//         merged = append(merged, nums2[i])
//     }
//     sort.Ints(merged)
//     midIdx := len(merged) / 2
//     if len(merged) % 2 != 0 {
//         // odd array len, return the mid
//         return float64(merged[midIdx])
//     }
//     var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
//     return avg
// }


/*
    approach : merge using two pointers so no need to sort, and then grab middle
    - merge both arrays into a new m+n array using 2 pointers
    - if len(newArray) is even, grab the middle and return
    - or else grab the 2 middle elements and return the avg between them (even len array)
    
    time: o(m+n)
    space: o(1)
*/
// func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//     merged := []int{}
//     n1 := 0
//     n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] < nums2[n2] {
//             merged = append(merged, nums1[n1])
//             n1++
//         } else {
//             merged = append(merged, nums2[n2])
//             n2++
//         }
//     }
//     if n1 < len(nums1) {
//         merged = append(merged, nums1[n1:]...)
//     }
//     if n2 < len(nums2) {
//         merged = append(merged, nums2[n2:]...)
//     }
    
//     midIdx := len(merged) / 2
//     if len(merged) % 2 != 0 {
//         // odd array len, return the mid
//         return float64(merged[midIdx])
//     }
//     var avg float64 = float64(merged[midIdx]+merged[midIdx-1])/2.0 
//     return avg
// }



