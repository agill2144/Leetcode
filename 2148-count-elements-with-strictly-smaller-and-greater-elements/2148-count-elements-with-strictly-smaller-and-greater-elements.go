func countElements(nums []int) int {
    if nums == nil || len(nums) == 0 || len(nums) == 1 {return 0}
    sort.Ints(nums)
    min := nums[0]
    max := nums[len(nums)-1]
    count := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] > min && nums[i] < max {count++}
    }
    return count
}