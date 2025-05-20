func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    out := [][]int{}
    for i := 0; i < n; i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        left := i+1
        right := n-1
        for left < right {
            sum := nums[i] + nums[left] + nums[right]
            if sum == 0 {
                out = append(out, []int{nums[i], nums[left], nums[right]})
                left++
                for left < right && nums[left] == nums[left-1] {left++}
                right--
                for left < right && nums[right] == nums[right+1] {right--}
            } else if sum > 0 {
                right--
            } else {
                left++
            }
        }
    }
    return out
}
// n^2 - two sum compliment search using hashing
// func threeSum(nums []int) [][]int {
//     target := 0
//     n := len(nums)
//     set := map[[3]int]bool{}
//     out := [][]int{}
//     for i := 0; i < n; i++ {
//         iTarget := target-nums[i]
//         cp := map[int]bool{}
//         for j := i+1; j < n; j++ {
//             diff := iTarget - nums[j]
//             if cp[diff] {
//                 tmp := []int{nums[i], nums[j], diff}
//                 sort.Ints(tmp)
//                 if !set[[3]int{tmp[0], tmp[1], tmp[2]}] {
//                     set[[3]int{tmp[0], tmp[1], tmp[2]}] = true
//                     out = append(out, tmp)
//                 }
//             }
//             cp[nums[j]] = true
//         }
//     }
//     return out
// }
// brute force: n^3
// func threeSum(nums []int) [][]int {
//     target := 0
//     n := len(nums)
//     set := map[[3]int]bool{}
//     out := [][]int{}
//     for i := 0; i < n; i++ {
//         for j := i+1; j < n; j++ {
//             for k := j+1; k < n; k++ {
//                 sum := nums[i] + nums[j] + nums[k]
//                 if sum == target {
//                     tmp := []int{nums[i], nums[j], nums[k]}
//                     sort.Ints(tmp)
//                     if !set[[3]int{tmp[0], tmp[1], tmp[2]}] {
//                         set[[3]int{tmp[0], tmp[1], tmp[2]}] = true
//                         out = append(out, tmp)
//                     }
//                 }
//             }
//         }
//     }
//     return out
// }
