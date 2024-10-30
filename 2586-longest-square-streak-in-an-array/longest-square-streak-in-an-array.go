func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    maxStreak := 0
    for i := 0; i < len(nums); i++ {
        curr := nums[i]*nums[i]
        count := 1
        for search(nums, i, curr) {
            count++
            curr *= curr
        }
        maxStreak = max(maxStreak, count)
    }
    if maxStreak == 1 {return -1}
    return maxStreak
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
//     set := map[int]bool{}
//     start := math.MaxInt64
//     end := math.MinInt64
//     for i := 0; i < len(nums); i++ {
//         set[nums[i]] = true
//         start = min(start, nums[i])
//         end = max(end, nums[i])
//     }
//     maxStreak := 1
//     for i := 0; i < len(nums); i++ {
//         curr := nums[i]
//         count := 0
//         for set[curr] && curr < int(math.Pow(10.0, 5)) {
//             curr = curr * curr
//             count++
//         }
//         maxStreak = max(maxStreak, count)
//     }
//     if maxStreak == 1 {return -1}
//     return maxStreak
// }