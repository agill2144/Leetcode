func minIncrementForUnique(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}
    sort.Ints(nums)
    newPrev := nums[0]
    count := 0
    for i := 1; i < n; i++ {
        if nums[i] <= newPrev {
            newCurr := newPrev+1
            diff := newCurr - nums[i]
            if diff == 0 {
                count++
            } else {
                count+=diff
            }
            
            newPrev = newCurr
        } else {
            newPrev = nums[i]
        }
    }
    return count
}