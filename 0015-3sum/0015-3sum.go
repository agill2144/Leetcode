func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    set := map[[3]int]struct{}{}
    out := [][]int{}
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {continue}
        target := 0-nums[i]
        left := i+1
        right := len(nums)-1
        for left < right {
            total := nums[left] + nums[right]
            if total == target {
                tmp := [3]int{nums[i],nums[left], nums[right]}
                if _, ok := set[tmp]; !ok {
                    set[tmp] = struct{}{}
                    out = append(out, tmp[:])
                }
                left++
                for left < right && nums[left] == nums[left-1] {left++}
                right--
                for right > left && nums[right] == nums[right+1] {right--}
            } else if total > target {
                right--
            } else {
                left++
            }
        }
    }
    return out
}
// func threeSum(nums []int) [][]int {
//     sort.Ints(nums)
//     set := map[[3]int]struct{}{}
//     out := [][]int{}
//     for i := 0; i < len(nums); i++ {
//         target := 0 - nums[i]
//         if i != 0 && nums[i] == nums[i-1] {continue}
//         for j := i+1; j < len(nums); j++ {
//             diff := target - nums[j]
//             if diff < nums[i] {continue}
//             found := binarySearch(diff, j+1, nums)
//             if found {
//                 tmp := [3]int{nums[i],nums[j], diff}
//                 if _, ok := set[tmp]; !ok {
//                     set[tmp] = struct{}{}
//                     out = append(out, tmp[:])
//                 }
//             }
//         }
//     }
//     return out
// }

// func binarySearch(target, left int, nums []int) bool {
//     right := len(nums)-1
//     for left <= right {
//         mid := left + (right-left) / 2
//         if nums[mid] == target {
//             return true
//         } else if target > nums[mid] {
//             left = mid+1
//         } else {
//             right = mid-1
//         }
//     }
//     return false
// }

// func threeSum(nums []int) [][]int {
//     set := map[[3]int]struct{}{}
//     out := [][]int{}
//     for i := 0; i < len(nums); i++ { // o(n)
//         target := 0-nums[i]
//         numsSet := map[int]struct{}{}
//         for j := i+1; j < len(nums); j++ { // o(n)
//             newTarget := target-nums[j]
//             if _, ok := numsSet[newTarget]; ok {
//                 tmp := []int{nums[i], nums[j], newTarget} 
//                 sort.Ints(tmp) // o(3log3)
//                 if _, ok := set[[3]int{tmp[0],tmp[1],tmp[2]}]; !ok {
//                     set[[3]int{tmp[0],tmp[1],tmp[2]}] = struct{}{}
//                     out = append(out, tmp)
//                 }
//             }
//             numsSet[nums[j]] = struct{}{}
//         }
//     }
//     return out
// }


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