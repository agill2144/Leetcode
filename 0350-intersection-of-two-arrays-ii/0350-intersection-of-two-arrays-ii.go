func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    
    out := []int{}
    n1 := 0
    n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {
            out = append(out, nums1[n1])
            n1++
            n2++
        } else if nums1[n1] < nums2[n2] {
             // [ 1 ..
            //  [ 2 ..
            n1++
        } else {
            n2++
        }
    }
    return out
}