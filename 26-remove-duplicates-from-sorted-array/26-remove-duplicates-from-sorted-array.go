func removeDuplicates(nums []int) int {
    // points to an index that should have the next uniq element
    slow := 0

    for i := 0; i < len(nums); i++ {
        if i == 0 {
            slow++
            continue
        }
        
        if nums[i] == nums[i-1] {
            continue
        } else {
            nums[slow] = nums[i]
            slow++
        }
    }
    return slow
}