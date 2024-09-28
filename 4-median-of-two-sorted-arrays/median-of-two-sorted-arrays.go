func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1) + len(nums2) 
    ele1 := -1
    idx1 := (n/2)-1
    ele2 := -1
    idx2 := n/2
    idx := 0
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Val := nums1[n1]
        n2Val := nums2[n2]
        if n1Val < n2Val {
            if idx == idx1 {
                ele1 = n1Val
            } else if idx == idx2 {
                ele2 = n1Val
            }
            n1++
        } else {
            if idx == idx1 {
                ele1 = n2Val
            } else if idx == idx2 {
                ele2 = n2Val
            }
            n2++
        }
        idx++
    }
    if ele1 == -1 || ele2 == -1  {
        for n1 < len(nums1) {
            n1Val := nums1[n1]
            if idx == idx1 {
                ele1 = n1Val
            } else if idx == idx2 {
                ele2 = n1Val
            }
            n1++
            idx++
        }

        for n2 < len(nums2) {
            n2Val := nums2[n2]
            if idx == idx1 {
                ele1 = n2Val
            } else if idx == idx2 {
                ele2 = n2Val
            }
            n2++
            idx++
        }
    }
    if n % 2 == 0 {
        return (float64(ele1) + float64(ele2)) / 2.0
    }
    return float64(ele2)
}