func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    left := 0
    for i := 0; i < len(nums1); i++ {
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