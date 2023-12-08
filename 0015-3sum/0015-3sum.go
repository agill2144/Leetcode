func threeSum(nums []int) [][]int {
    set := map[[3]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        target := 0-nums[i]
        numsSet := map[int]struct{}{}
        for j := i+1; j < len(nums); j++ {
            newTarget := target-nums[j]
            if _, ok := numsSet[newTarget]; ok {
                tmp := [3]int{nums[i], nums[j], newTarget}
                sort.Ints(tmp[:])
                if _, ok := set[tmp]; !ok {
                    set[tmp] = struct{}{}
                    out = append(out, tmp[:])
                }
            }
            numsSet[nums[j]] = struct{}{}
        }
    }
    return out
}


// func threeSum(nums []int) [][]int {
//     set := map[[3]int]struct{}{}
//     out := [][]int{}
//     for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             for k := j+1; k < len(nums); k++ {
//                 if nums[i] + nums[j] + nums[k] == 0 {
//                     tmp := [3]int{nums[i], nums[j], nums[k]}
//                     sort.Ints(tmp[:])
//                     if _, ok := set[tmp]; !ok {
//                         set[tmp] = struct{}{}
//                         out = append(out, []int{tmp[0], tmp[1], tmp[2]})
//                     }
//                 }
//             }
//         }
//     }
//     return out
// }