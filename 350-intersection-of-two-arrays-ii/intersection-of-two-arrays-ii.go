func intersect(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersect(nums2, nums1)}
    sort.Ints(nums1)
    sort.Ints(nums2)
    // loop over smaller array and binary search on larger array
    // nums1 is always smaller
    out := []int{}
    left := 0
    for i := 0; i < len(nums1); i++ {
        idx := search(nums2, left, nums1[i])
        if idx != -1 {
            out = append(out, nums1[i])
            left = idx+1
        }   
    }
    return out
}

func search(nums []int, left, target int) int {
    ans := -1
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            ans = mid
            right = mid-1
            continue
        }
        if target > nums[mid] {left=mid+1} else {right=mid-1}
    }
    return ans
}