// time = o(log min(n1, n2))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n1 := len(nums1)
    n2 := len(nums2)
    if n2 < n1 {return findMedianSortedArrays(nums2, nums1)}
    total := n1+n2
    half := total/2
    left := 0
    right := n1
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid < n1 {r1 = nums1[mid]}
        if mid2 < n2 {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1, l2)) + float64(min(r1, r2))) / 2.0
            }
            return float64(min(r1, r2))
        }
        if l2 > r1 {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return 0.0
}
