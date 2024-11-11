func removeDuplicates(nums []int) int {
    slow := 1
    for i := 1; i < len(nums); i++ {
        curr := nums[i]; prev := nums[i-1]
        if curr != prev {
            nums[slow] = nums[i]
            slow++
        }
    }
    return slow
}