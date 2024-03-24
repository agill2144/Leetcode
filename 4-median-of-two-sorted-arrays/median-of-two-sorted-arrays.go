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

*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    sortedArr := []int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] < nums2[n2] {
            sortedArr = append(sortedArr, nums1[n1])
            n1++
        } else {
            sortedArr = append(sortedArr, nums2[n2])
            n2++
        }
    }
    for n1 < len(nums1) {sortedArr = append(sortedArr, nums1[n1]); n1++}
    for n2 < len(nums2) {sortedArr = append(sortedArr, nums2[n2]); n2++}
    mid := (len(sortedArr)-1) / 2
    if len(sortedArr) % 2 != 0 {
        return float64(sortedArr[mid])
    }
    return (float64(sortedArr[mid]) + float64(sortedArr[mid+1])) / 2.0    
}
