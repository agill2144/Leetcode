func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {    
    total := len(nums1) + len(nums2)
    mid1 := (total-1)/2
    ele1 := -1
    mid2 := mid1+1
    ele2 := -1
    ptr := 0
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Val := nums1[n1]
        n2Val := nums2[n2]
        if n1Val < n2Val {
            if ptr == mid1 {ele1=n1Val}
            if ptr == mid2 {ele2=n1Val}
            n1++
        } else {
            if ptr == mid1 {ele1=n2Val}
            if ptr == mid2 {ele2=n2Val}
            n2++
        }
        ptr++
    }
    for n1 < len(nums1){
        n1Val := nums1[n1]
        if ptr == mid1 {ele1 = n1Val}
        if ptr == mid2 {ele2 = n1Val}
        ptr++
        n1++
    }
    for n2 < len(nums2){
        n2Val := nums2[n2]
        if ptr == mid1 {ele1 = n2Val}
        if ptr == mid2 {ele2 = n2Val}
        ptr++
        n2++
    }
    if total % 2 == 0 {
        return (float64(ele1) + float64(ele2)) / 2.0
    }
    return float64(ele1)
}