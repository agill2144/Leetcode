// sorted?
// binary search
// tc = o(nlogn)
// sc = o(1)
func twoSum(nums []int, target int) []int {
    n := len(nums)
    for i := 0; i < n; i++ {
        left := i+1
        right := n-1
        for left <= right {
            mid := left+(right-left)/2
            sum := nums[i] + nums[mid]
            if sum == target {return []int{i+1, mid+1}}
            if sum > target {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return nil
}
// sorted?
// two ptrs
// tc = o(n)
// sc = o(1)
// func twoSum(nums []int, target int) []int {
//     n := len(nums)
//     if n <= 1 {return nil}
//     left := 0
//     right := n-1

//     for left < right {
//         sum := nums[left] + nums[right]
//         if sum == target {return []int{left+1, right+1}}
//         if sum > target {
//             right--
//         } else {
//             left++
//         }
//     }
//     return nil
// }