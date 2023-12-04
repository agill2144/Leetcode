func removeDuplicates(nums []int) int {
    slow := 1
    fast := 1
    // fast ptr checks whether it has a uniq element
    // if yes, it gives it to slow ptr, slow ptr swaps with fast ptr
    // and slow ptr moves forward
    
    for fast < len(nums) {
        if nums[fast] != nums[fast-1] {
            // uniq
            nums[slow] = nums[fast]
            slow++
        }
        fast++
    }
    return slow
}