func findPeakElement(nums []int) int {
    n := len(nums)
    left := 0
    right := n-1
    for left <= right {
        mid := left + (right-left)/2
        if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == n-1 || nums[mid] > nums[mid+1]) {
            return mid
        }
        if (mid == n-1 || nums[mid+1] > nums[mid]) {
            left = mid+1
        } else {
            right = mid-1
        }
    }
    return -1
}
// brute force: linear scan
// func findPeakElement(nums []int) int {
//     // strictly greater than its neighbors
//     // not >= , strictly > 
//     for i := 0; i < len(nums); i++ {
//         prev := math.MinInt64
//         if i-1 >= 0 {prev = nums[i-1]}
//         curr := nums[i]
//         next := math.MinInt64
//         if i+1 < len(nums) {next = nums[i+1]}
//         if curr > prev && curr > next {return i}
//     }
//     return -1
// }