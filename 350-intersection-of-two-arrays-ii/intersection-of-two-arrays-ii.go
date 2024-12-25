// if sorted, two ptrs or binary search
// approach: binary search on larger array while looping over smaller array
// m = len of smaller array
// n = len of larger array
// tc = m * logn
// sc = o(1)
func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    // nums1 is always smaller
    left := 0
    for i := 0; i < len(nums1); i++ {
        target := nums1[i]
        idx := search(target, left, nums2)
        if idx != -1 {
            out = append(out, nums1[i])
            left = idx+1
        }
    }    
    return out
}

// search for left most pos of target
// why?
// because its possible there are multiple instances of target that are needed
// and if we returned a position thats the right most
// we end up excluding the same target that went unused
func search(target, left int, nums []int) int {
    idx := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if target <= nums[mid] {
            if nums[mid] == target {idx = mid}
            right = mid-1
        } else if target > nums[mid] {
            left = mid+1
        }
    }
    return idx
}

// if sorted, two ptrs or binary search
// approach: two ptrs
// func intersect(nums1 []int, nums2 []int) []int {
//     sort.Ints(nums1)
//     sort.Ints(nums2)
//     out := []int{}
//     n1 := 0; n2 := 0
//     for n1 < len(nums1) && n2 < len(nums2) {
//         if nums1[n1] == nums2[n2] {
//             out = append(out, nums1[n1])
//             n1++; n2++
//         } else if nums1[n1] < nums2[n2] {
//             n1++
//         } else {
//             n2++
//         }
//     }
//     return out
// }
// approach : put smaller array into freq map ( because we have to loop over both array regardless! )
// func intersect(nums1 []int, nums2 []int) []int {
//     if len(nums2) < len(nums2) {return intersect(nums2, nums1)} 
//     freq := map[int]int{}
//     for i := 0; i < len(nums1); i++ {
//         freq[nums1[i]]++
//     }
//     out := []int{}
//     for i := 0; i < len(nums2); i++ {
//         if freq[nums2[i]] > 0 {
//             out = append(out, nums2[i])
//             freq[nums2[i]]--
//         }
//     }
//     return out
// }