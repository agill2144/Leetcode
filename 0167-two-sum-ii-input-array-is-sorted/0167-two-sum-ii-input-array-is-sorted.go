func twoSum(nums []int, target int) []int {
    left := 0
    right := len(nums)-1
    
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