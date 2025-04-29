func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        left := i+1
        right := len(nums)-1
        for left < right {
            total := nums[i] + nums[left] + nums[right]
            if total == 0 {
                out = append(out, []int{nums[i], nums[left], nums[right]})
                left++ 
                for left < right && nums[left] == nums[left-1] {
                    left++
                }
                right--
                for left < right && nums[right] == nums[right+1] {
                    right--
                }
            } else if total > 0 {
                right--
            } else {
                left++
            }
        }
    }
    return out

}
// func threeSum(nums []int) [][]int {
//     target := 0
//     set := map[[3]int]bool{}
//     n := len(nums)
//     out := [][]int{}
//     for i := 0; i < n; i++ {
//         iTarget := target - nums[i]
//         set2 := map[int]bool{}
//         for j := i+1; j < n; j++ {
//             diff := iTarget - nums[j]
//             if set2[diff] {
//                 tmp := []int{nums[i], nums[j], diff}
//                 sort.Ints(tmp)
//                 if !set[[3]int{tmp[0],tmp[1],tmp[2]}] {
//                     out = append(out, tmp)
//                     set[[3]int{tmp[0],tmp[1],tmp[2]}] = true
//                 }
//             }
//             set2[nums[j]] = true
//         }
//     }
//     return out
// }