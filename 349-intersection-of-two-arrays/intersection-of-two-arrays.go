func intersection(nums1 []int, nums2 []int) []int {
    if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
    // nums1 is always smaller

    // time = o(n1)
    // space = o(n1)
    n1Map := map[int]bool{}
    for i := 0; i < len(nums1); i++ {
        n1Map[nums1[i]] = false // false = not added to our output yet
    }

    // time = o(n2)
    out := []int{}
    for i := 0; i < len(nums2); i++ {
        num := nums2[i]
        if used, ok := n1Map[num]; ok && !used {
            out = append(out, num)
            n1Map[num] = true // used, added to output, cannot be re-used
        }
    }
    return out

    // time = o(n1) + o(n2) + o(n2) = o(n1+n2)
    // space = o(n1) + o(n2)
}


// hash hash
// n1 = len(nums1) 
// n2 = len(nums2)
// func intersection(nums1 []int, nums2 []int) []int {
//     if len(nums2) < len(nums1) {return intersection(nums2, nums1)}
//     // nums1 is always smaller

//     // time = o(n1)
//     // space = o(n1)
//     n1Set := map[int]struct{}{}
//     for i := 0; i < len(nums1); i++ {
//         n1Set[nums1[i]] = struct{}{}
//     }

//     // time = o(n2)
//     // space = o(n2) ; worst case n1 == n2 ( identical elements )
//     outSet := map[int]struct{}{}
//     for i := 0; i < len(nums2); i++ {
//         num := nums2[i]
//         if _, ok := outSet[num]; ok {continue}
//         if _, ok := n1Set[num]; ok {
//             outSet[num] = struct{}{}
//         }
//     }
//     out := []int{}
//     // time = o(n2)
//     for k, _ := range outSet {
//         out = append(out, k)
//     }
//     return out

//     // time = o(n1) + o(n2) + o(n2) = o(n1+n2)
//     // space = o(n1) + o(n2)
// }