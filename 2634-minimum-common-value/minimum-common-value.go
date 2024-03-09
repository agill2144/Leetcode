func getCommon(nums1 []int, nums2 []int) int {
    n1, n2 := 0,0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {return nums1[n1]}
        if nums1[n1] > nums2[n2] {
            n2++
        } else {
            n1++
        }
    }
    return -1
}