func minIncrementForUnique(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}
    sort.Ints(nums)
    newPrev := nums[0]
    count := 0
    for i := 1; i < n; i++ {
        if nums[i] <= newPrev {
            newCurr := newPrev+1
            count += max(1, newCurr - nums[i])
            newPrev = newCurr
        } else {
            newPrev = nums[i]
        }
    }
    return count
}