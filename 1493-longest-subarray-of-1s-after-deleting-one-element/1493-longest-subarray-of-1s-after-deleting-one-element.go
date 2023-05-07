func longestSubarray(nums []int) int {
    count := 0
    maxSize := 0
    left := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] == 0 {count++}
        for count > 1 {
            if nums[left] == 0 {count--}
            left++
        }
        if i-left > maxSize {maxSize = i-left}
    }
    return maxSize
}