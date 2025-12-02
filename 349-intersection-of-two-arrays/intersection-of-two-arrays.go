func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    left := 0
    for i := 0; i < len(nums1); i++ {
        // if dupes were allowed to be collected, we would drop this condition
        // and we would want the binary search to return to the leftMost idx of an element
        // such that if more dupes of next element exists in nums2, we can also collect them
        // only when dupes are allowed*
        // in this problem dupes are not allowed
        if i > 0 && nums1[i] == nums1[i-1] {continue}
        target := nums1[i]
        idx := search(nums2, target, left)
        if idx != -1 {
            out = append(out, nums1[i])
            left = idx+1
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
            left = mid+1
        } else if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return idx
}