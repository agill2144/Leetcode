func removeElement(nums []int, val int) int {
    
    slow := 0    
    for i := 0; i < len(nums); i++ {
        if nums[i] == val {
            continue
        }
        nums[slow] = nums[i]
        slow++
    }
    return slow
}