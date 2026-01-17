func threeSumSmaller(nums []int, target int) int {
    count := 0
    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
        // no need to escape if i val is same as prev
        // because the question says just count 
        // does not say count unique
        left := i+1
        right := len(nums)-1
        for left < right {
            total := nums[i] + nums[left] + nums[right]
            if total < target {
                count += (right-left)
                left++
            } else {
                right--
            }
        }
    }
    return count
}

// brute force
// func threeSumSmaller(nums []int, target int) int {
//     count := 0
//    for i := 0; i < len(nums); i++ {
//         for j := i+1; j < len(nums); j++ {
//             for k := j+1; k < len(nums); k++ {
//                 if nums[i] + nums[j] + nums[k] < target {
//                     count++
//                 }
//             }
//         }
//     }
//     return count
// }