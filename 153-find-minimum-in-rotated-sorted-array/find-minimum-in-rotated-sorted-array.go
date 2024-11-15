func findMin(nums []int) int {
    for i := 1; i < len(nums); i++ {
        if nums[i] < nums[i-1] {return nums[i]}
    }
    return nums[0]
}