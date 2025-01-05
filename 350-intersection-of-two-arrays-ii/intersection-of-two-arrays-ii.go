func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    p1 := 0
    p2 := 0
    for p1 < len(nums1) && p2 < len(nums2) {
        if nums1[p1] == nums2[p2] {
            out = append(out, nums1[p1])
            p1++
            p2++
        } else if nums1[p1] < nums2[p2] {
            p1++
        } else {
            p2++
        }
    }
    return out
}