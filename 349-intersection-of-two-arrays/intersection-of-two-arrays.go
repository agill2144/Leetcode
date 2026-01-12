func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1){return intersection(nums2, nums1)}  
    // nums1 is always smaller in size
    sort.Ints(nums1)
    sort.Ints(nums2)
    left := 0
    out := []int{}
    for i := 0; i < len(nums1); i++ {
        if i != 0 && nums1[i] == nums1[i-1] {continue}
        target := nums1[i]
        idx := search(left, target, nums2)
        if idx != -1 {
            out = append(out, nums1[i])
            left = idx+1
        }
    }
    return out
}

func search(left, target int, nums []int) int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            idx = mid
            left = mid+1
        } else if target < nums[mid] {
            right = mid-1
        } else {
            left = mid+1
        }
    }
    return idx
}
// func intersection(nums1 []int, nums2 []int) []int {
//     sort.Ints(nums1)
//     sort.Ints(nums2)
//     n1, n2 := 0, 0
// 	out := []int{}
// 	for n1 < len(nums1) && n2 < len(nums2) {
// 		n1Val := nums1[n1]
// 		n2Val := nums2[n2]
// 		if (n1Val == n2Val) {
//             if len(out) == 0 || n1Val != out[len(out)-1] {
//     			out = append(out, n1Val)
//             }
// 			n1++
// 			n2++
// 		} else if n1Val < n2Val {
// 			n1++
// 		} else {
// 			n2++
// 		}
// 	}
// 	return out
// }
