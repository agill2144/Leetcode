func threeSum(nums []int) [][]int {
    set := map[[3]int]bool{}
    res := [][]int{}
    for i := 0; i < len(nums); i++ {
        iTarget := 0-nums[i]
        m := map[int]bool{}
        for j := i+1; j < len(nums); j++ {
            diff := iTarget - nums[j]
            if m[diff] {
                tmp := []int{nums[i], nums[j], diff}
                sort.Ints(tmp)
                if !set[[3]int{tmp[0], tmp[1], tmp[2]}] {
                    set[[3]int{tmp[0], tmp[1], tmp[2]}] = true
                    res = append(res, tmp)
                }   
            }
            m[nums[j]] = true
        }
    }
    return res
}
// brute force, 3 nested loops, and check each triplet sum
// tc = o(n * n * (n * 3log3 + 3) = o(n^3)
// sc = o(n)
// func threeSum(nums []int) [][]int {
//     set := map[[3]int]bool{}
//     res := [][]int{}
//     n := len(nums)
//     for i := 0; i < n-2; i++ {
//         for j := i+1; j < n-1; j++ {
//             for k := j+1; k < n; k++ {
//                 sum := nums[i] + nums[j] + nums[k]
//                 if sum == 0 {
//                     tmp := []int{nums[i], nums[j], nums[k]}
//                     sort.Ints(tmp)
//                     if !set[[3]int{tmp[0], tmp[1], tmp[2]}] {
//                         set[[3]int{tmp[0], tmp[1], tmp[2]}] = true
//                         res = append(res, tmp)
//                     }
//                 }
//             }
//         }
//     } 
//     return res
// }