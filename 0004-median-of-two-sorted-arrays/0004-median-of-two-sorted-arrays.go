func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums2) < len(nums1) {return findMedianSortedArrays(nums2, nums1)}
    // find the split using binary search
    
    n1 := len(nums1)
    n2 := len(nums2)
    left := 0
    right := n1
    
    for left <= right {
        partA := left + (right-left)/2
        partB := (n1+n2)/2 - partA
        fmt.Println(nums1, nums2, partA, partB)
        l1 := math.MinInt64
        if partA != 0 {l1 = nums1[partA-1]}
        l2 := math.MinInt64
        if partB != 0 {l2 = nums2[partB-1]}
        
        r1 := math.MaxInt64
        if partA < n1 {r1 = nums1[partA]}
        r2 := math.MaxInt64
        if partB < n2 {r2 = nums2[partB]}
        
        if l1 <= r2 && l2 <= r1 {
            if (n1+n2) % 2 == 0 {
                return (math.Max(float64(l1), float64(l2)) + math.Min(float64(r1), float64(r2))) / 2.0 
            } else {
                return math.Min(float64(r1), float64(r2))
            }
        } else if l2 > r1 {
            left = partA+1
        } else {
            right = partA-1
        }    
    }
    return -1
}