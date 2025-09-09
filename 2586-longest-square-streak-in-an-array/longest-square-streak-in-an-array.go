func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    maxCount := 0
    for i := 0; i < len(nums); i++ {
        count := 1
        target := nums[i]*nums[i]
        for search(nums,i+1,target) {
            count++
            target *= target
        }
        maxCount = max(maxCount, count)
    }

    if maxCount == 1 {maxCount = -1}
    return maxCount
}

func search(nums []int, left, target int) bool {
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return true
        }
        if target > nums[mid] {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return false
}

// hashing: longest consec seq pattern
// tc = o(n) + o(n * log n)
// sc = o(n)
// func longestSquareStreak(nums []int) int {
//     set := map[float64]bool{}
//     for i := 0; i < len(nums); i++ {set[float64(nums[i])] = true}
//     maxStreak := 0
//     for i := 0; i < len(nums); i++ {
//         sqrt := math.Sqrt(float64(nums[i]))
//         if set[sqrt] {continue}
//         start := float64(nums[i])
//         count := 0
//         for set[start] {
//             count++
//             start *= start
//         }
//         maxStreak = max(maxStreak, count)
//     }
//     if maxStreak == 1 {maxStreak = -1}
//     return maxStreak
// }