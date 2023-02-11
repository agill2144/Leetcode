func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1) }
    
    n1 := len(nums1)
    n2 := len(nums2)
    
    left := 0
    right := n1
    
    for left <= right {
        partX := left + (right-left)/2
        partY := (n1+n2)/2 - partX
        
        l1 := math.MinInt64
        if partX != 0 {l1 = nums1[partX-1]}
        l2 := math.MinInt64
        if partY != 0 {l2 = nums2[partY-1]}
        
        r1 := math.MaxInt64
        if partX < n1 {r1 = nums1[partX]}
        r2 := math.MaxInt64
        if partY < n2 {r2 = nums2[partY]}
        
        if l1 <= r2 && l2 <= r1 {
            if (n1+n2) % 2 == 0 {
                return (math.Min(float64(r1), float64(r2)) + math.Max(float64(l1), float64(l2))) / 2.0
            } else {
                return math.Min(float64(r1), float64(r2))
            }
        } else if l1 > r2 {
            right = partX-1
        } else {
            left = partX+1
        }
    }
    return -1
}   