func addedInteger(nums1 []int, nums2 []int) int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    return nums2[0]-nums1[0]
}

// func addedInteger(nums1 []int, nums2 []int) int {
//     n1 := len(nums1)
//     n2 := len(nums2)
//     if n2 != n1 {return math.MaxInt64}
    
//     freq := map[int]int{}
//     for i := 0; i < len(nums2); i++ {freq[nums2[i]]++}

//     f := len(freq)
//     for i := -1000; i <= 1000; i++ {
//         count := 0
//         local := map[int]int{}
//         for j := 0; j < n1; j++ {
//             val := nums1[j]+i
//             desiredCount, ok := freq[val]
//             currentCount := local[val]
//             if !ok || currentCount == desiredCount {
//                 break
//             }

//             local[val]++
//             if local[val] == desiredCount {
//                 count++
//                 if count == f {
//                     return i
//                 }
//             }
            
//         }
//     }
//     return math.MaxInt64
// }