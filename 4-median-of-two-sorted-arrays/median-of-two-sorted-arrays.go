func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    idx := 0
    n1 := len(nums1)
    n2 := len(nums2)
    total := n1+n2
    mid := (total-1)/2
    p1, p2 := 0,0
    ele1, ele2 := math.MinInt64, math.MinInt64

    for p1 < n1 || p2 < n2 {
        p1Val := math.MaxInt64
        if p1 < n1 {p1Val = nums1[p1]}
        p2Val := math.MaxInt64
        if p2 < n2 {p2Val = nums2[p2]}

        if p1Val < p2Val {
            if idx == mid {
                ele1 = p1Val
            }  else if idx == mid+1 {
                ele2 = p1Val
            }
            p1++
            idx++
        } else {
            if idx == mid {
                ele1 = p2Val
            }  else if idx == mid+1 {
                ele2 = p2Val
            }
            p2++
            idx++
        }
    }

    if total % 2 == 0 {
        return (float64(ele1)+float64(ele2))/2.0
    }
    return float64(ele1)
}