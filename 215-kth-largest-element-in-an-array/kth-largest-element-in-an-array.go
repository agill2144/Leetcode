func findKthLargest(nums []int, k int) int {
    freq := map[int]int{}
    minVal := math.MaxInt64
    maxVal := math.MinInt64
    for i := 0; i < len(nums); i++ {
        freq[nums[i]]++
        minVal = min(minVal, nums[i])
        maxVal = max(maxVal, nums[i])
    }
    idx := -1
    for i := minVal; i <= maxVal; i++ {
        idx += freq[i]
        if idx >= len(nums)-k {return i}
    }
    return -1
}

// quick select ( sort the half we care about )
// avg time = o(n)
// space = o(1)
// func findKthLargest(nums []int, k int) int {
//     n := len(nums)
//     targetIdx := n-k
//     left := 0
//     right := n-1
//     for left <= right {
//         pivot := right
//         ns := left
//         for i := left; i < pivot; i++ {
//             if nums[i] <= nums[pivot] {
//                 nums[i], nums[ns] = nums[ns], nums[i]
//                 ns++
//             }
//         }
//         nums[ns], nums[pivot] = nums[pivot], nums[ns]
//         if ns == targetIdx {
//             return nums[ns]
//         }
//         if targetIdx > ns {
//             left = ns+1
//         } else {
//             right = ns-1
//         }
//     }
//     return -1
// }
