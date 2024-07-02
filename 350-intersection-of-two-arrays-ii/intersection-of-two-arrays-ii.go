func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    n1, n2 := 0, 0
    out := []int{}
    for n1 < len(nums1) && n2 < len(nums2) {
        n1Val := nums1[n1]
        n2Val := nums2[n2]
        if n1Val == n2Val {
            out = append(out, n1Val)
            n1++
            n2++
        } else if n1Val < n2Val {
            n1++
        } else {
            n2++
        }
    }
    return out
}

// func intersect(nums1 []int, nums2 []int) []int {
//     n1Freq := map[int]int{}
//     for i := 0; i < len(nums1); i++ {
//         n1Freq[nums1[i]]++
//     }
//     out := []int{}
//     for i := 0; i < len(nums2); i++ {
//         if val := n1Freq[nums2[i]]; val > 0 {
//             out = append(out, nums2[i])
//             n1Freq[nums2[i]]--
//         }
//     }
//     return out
// }