func findKthLargest(nums []int, k int) int {
    var dfs func(l, r int) (int, bool)
    dfs = func(l, r int) (int, bool) {
        // base
        // if l >= r {return -1, false}
        
        // logic
        pivotIdx := r
        nextSmaller := l
        for i := l; i < pivotIdx; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[i], nums[nextSmaller] = nums[nextSmaller], nums[i]
                nextSmaller++
            }
        }
        nums[nextSmaller], nums[pivotIdx] = nums[pivotIdx], nums[nextSmaller]
        if len(nums) - k == nextSmaller {
            return nums[nextSmaller], true
        }
        if len(nums)-k > nextSmaller {
            if val, ok := dfs(nextSmaller+1, r); ok {
                return val, ok
            }
        }
        return dfs(l, nextSmaller-1)
    }
    val, _ := dfs(0, len(nums)-1)
    return val
    
}