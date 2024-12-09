func longestSquareStreak(nums []int) int {
    sort.Ints(nums)
    maxSize := -1
    for i := 0; i < len(nums); i++ {
        count := 1
        num := nums[i]
        for search(i+1, num*num, nums) {
            count++
            num *= num
        }
        maxSize = max(maxSize, count)
    }
    if maxSize == 1 {maxSize = -1}
    return maxSize
}
func search(left int, target int, nums []int) bool {
    if target > nums[len(nums)-1] {return false}
    right := len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {return true}
        if target < nums[mid] {right = mid-1} else { left = mid+1 }
    }
    return false
}