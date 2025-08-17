// sorted 
// 1. two ptrs
func twoSum(nums []int, target int) []int {
    n := len(nums)
    left := 0
    right := n-1
    for left < right {
        sum := nums[left] + nums[right]
        if sum == target {return []int{left+1, right+1}}
        if sum > target {
            right--
        } else {
            left++
        }
    }
    return nil
}

// compliment search will obv work
// tc = o(n)
// sc = o(n)
// func twoSum(nums []int, target int) []int {
//     idx := map[int]int{}
//     for i := 0; i < len(nums); i++ {
//         diff := target - nums[i]
//         val, ok := idx[diff]
//         if ok {return []int{val+1, i+1}}
//         idx[nums[i]] = i
//     }
//     return nil
// }