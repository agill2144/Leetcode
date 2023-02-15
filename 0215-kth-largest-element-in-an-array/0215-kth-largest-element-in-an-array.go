func findKthLargest(nums []int, k int) int {
    var dfs func(left, right int) int
    dfs = func(left, right int) int {        
        
        // base
        if left > right {return math.MinInt64}
        
        // logic
        ns := left
        pivotIdx := right
        for i := left; i < right; i++ {
            if nums[i] < nums[pivotIdx] {
                nums[ns], nums[i] = nums[i], nums[ns]
                ns++
            }
        }
        
        nums[ns], nums[pivotIdx] = nums[pivotIdx], nums[ns]
        if ns == len(nums)-k {
            return nums[ns]
        } else if len(nums)-k > ns {
            if val := dfs(ns+1, right); val != math.MinInt64 {return val}
        }
        return dfs(left, ns-1)
    }
    return dfs(0, len(nums)-1)
}