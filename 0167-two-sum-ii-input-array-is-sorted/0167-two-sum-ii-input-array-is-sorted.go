// two ptrs
// time: o(n)
// space: o(1)
// func twoSum(nums []int, target int) []int {
//     left := 0
//     right := len(nums)-1
    
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

func twoSum(nums []int, target int) []int{
    for i := 0; i < len(nums); i++ {
        num := nums[i]
        comp := target-num
        left := i+1
        right := len(nums)-1
        for left <= right {
            mid := left + (right-left)/2
            if nums[mid] == comp {
                return []int{i+1, mid+1}
            } else if nums[mid] > comp {
                right = mid-1
            } else {
                left = mid+1
            }
        }
    }
    return nil
}