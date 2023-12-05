func findMaxConsecutiveOnes(nums []int) int {
    slow := 0
    fast := 0
    max := 0
    for fast < len(nums) {
        if nums[fast] == 0 {
            // save the last window so far
            slow = fast+1
            fast++
        } else {
            if fast-slow+1 > max {
                max = fast-slow+1
            }            
            fast++
        }
    }
    return max
}