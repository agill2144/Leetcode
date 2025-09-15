func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    left := 0
    for i := 0; i < len(nums1); i++ {
        idx := search(nums2, nums1[i], left)
        if idx != -1 {
            left = idx+1
            out = append(out, nums1[i])
        }
    }
    return out
}

func search(nums []int, target int, left int) int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            idx = mid
            right = mid-1
        } else if target < nums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}
// two ptrs
// tc = o(n1 + n2) ; if we dont count sorting
// sc = o(1)
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