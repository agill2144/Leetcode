func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    out := []int{}
    n1 := 0; n2 := 0
    for n1 < len(nums1) && n2 < len(nums2) {
        if nums1[n1] == nums2[n2] {
            out = append(out, nums1[n1])
            n1++; n2++
        } else if nums1[n1] < nums2[n2] {
            n1++
        } else {
            n2++
        }
    }
    return out
}
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