func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1) + len(nums2)
    mid1 := (n-1)/2
    ele1 := -1
    mid2 := mid1+1
    ele2 := -1
    idx := 0
    n1 := 0; n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Val := nums1[n1]
        n2Val := nums2[n2]
        if n1Val < n2Val {
            if idx == mid1 {ele1 = n1Val}
            if idx == mid2 {ele2 = n1Val}
            idx++
            n1++
        } else {
            if idx == mid1 {ele1 = n2Val}
            if idx == mid2 {ele2 = n2Val}
            idx++
            n2++
        }
    }
    for n1 < len(nums1){
        n1Val := nums1[n1]
        if idx == mid1 {ele1 = n1Val}
        if idx == mid2 {ele2 = n1Val}
        idx++
        n1++
    }

    for n2 < len(nums2){
        n2Val := nums2[n2]
        if idx == mid1 {ele1 = n2Val}
        if idx == mid2 {ele2 = n2Val}
        idx++
        n2++
    }

    if n % 2 != 0 {return float64(ele1)}
    return (float64(ele1) + float64(ele2)) / 2.0

}