
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m := len(nums1)
    n := len(nums2)
    if n < m {return findMedianSortedArrays(nums2, nums1)}
    total := m+n
    left := 0
    right := m
    half := total/2
    for left <= right {
        mid := left + (right-left)/2
        mid2 := half-mid
        l1 := math.MinInt64
        l2 := math.MinInt64
        r1 := math.MaxInt64
        r2 := math.MaxInt64
        if mid-1 >= 0 {l1 = nums1[mid-1]}
        if mid2-1 >= 0 {l2 = nums2[mid2-1]}
        if mid < m {r1 = nums1[mid]}
        if mid2 < n {r2 = nums2[mid2]}

        if l1 <= r2 && l2 <= r1 {
            if total % 2 == 0 {
                return (float64(max(l1,l2)) + float64(min(r1,r2))) / 2.0
            }
            return float64(min(r1,r2))
        }
        if l1 > r2 {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return 0.0
}