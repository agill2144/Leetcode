func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    // nums1 is always the smaller array
    sort.Ints(nums1)
    sort.Ints(nums2)
    idx := 0
    out := []int{}
    for i := 0; i < len(nums1); i++ {
        target := nums1[i]
        left := idx
        right := len(nums2)-1
        ansIdx := -1
        for left <= right {
            mid := left + (right-left)/2
            if nums2[mid] >= target {
                if nums2[mid] == target {ansIdx = mid}
                right = mid-1
            } else {
                left = mid+1
            }
        }
        if ansIdx != -1 {
            out = append(out, nums1[i])
            idx = ansIdx+1
        }
    }
    return out
}
// func intersect(nums1 []int, nums2 []int) []int {
//     sort.Ints(nums1)
//     sort.Ints(nums2)
//     n1, n2 := 0, 0 
//     out := []int{}
//     for n1 < len(nums1) && n2 < len(nums2) {
//         n1Val := nums1[n1]
//         n2Val := nums2[n2]
//         if n1Val == n2Val {
//             out = append(out, n1Val)
//             n1++
//             n2++
//         } else if n1Val < n2Val {
//             n1++
//         } else {
//             n2++
//         }
//     }
//     return out
// }