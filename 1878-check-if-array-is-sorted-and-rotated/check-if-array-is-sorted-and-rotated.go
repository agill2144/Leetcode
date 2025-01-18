func check(nums []int) bool {
    n := len(nums)
    startIdx := -1
    for i := 0; i+1 < len(nums); i++ {
        if nums[i] > nums[i+1] {
            startIdx = i+1
            break
        }
    }
    if startIdx == -1 {return true}
    endIdx := startIdx-1
    if endIdx < 0 {endIdx = n-1}
    for (startIdx%n) != (endIdx%n) {
        if nums[startIdx%n] > nums[(startIdx+1)%n] {
            return false
        }
        startIdx++
    }
    return true
}   