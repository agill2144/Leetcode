// sorted?
// two ptrs
// binary search
func twoSum(nums []int, target int) []int {
    n := len(nums)
    if n <= 1 {return nil}
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