func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    ans := -1
    for i := 0; i < len(nums); i++ {
        count := 1
        target := nums[i]
        for search(nums, i+1, target*target) {
            count++
            target *= target
        }
        if count > 1 { ans = max(ans, count)}
    }
    return ans
}
func search(nums []int, left, target int) bool {
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return true}
        if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}
// func longestSquareStreak(nums []int) int {
//     set := map[float64]bool{}
//     for i := 0; i < len(nums); i++ {
//         set[float64(nums[i])] = true
//     }
//     ans := -1
//     for i := 0; i < len(nums); i++ {
//         sqrt := math.Sqrt(float64(nums[i]))
//         if !set[sqrt] {
//             count := 0
//             start := float64(nums[i])
//             for set[start] {
//                 count++
//                 start *= start
//             }
//             if count > 1 {ans = max(ans, count)}
//         }
//     }
//     return ans
// }