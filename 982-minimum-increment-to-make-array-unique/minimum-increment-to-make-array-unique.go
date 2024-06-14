func minIncrementForUnique(nums []int) int {
    n := len(nums)
    if n <= 1 {return 0}
    sort.Ints(nums)
    newNums := make([]int, n)
    copy(newNums, nums)
    count := 0
    for i := 1; i < n; i++ {
        if nums[i] <= newNums[i-1] {
            newPrev := newNums[i-1]
            newCurr := newPrev+1
            diff := newCurr - nums[i]
            if diff == 0 {
                count++
            } else {
                count+=diff
            }
            newNums[i] = newCurr
        }
    }
    return count
}