func sortArray(nums []int) []int {
    var dfs func(left, right int)
    dfs = func(left, right int) {
        // base 
        if left >= right {return}
        
        // logic
        mid := left + (right-left)/2
        dfs(left, mid)
        dfs(mid+1, right)
        
        l := left
        r := mid+1
        merged := []int{}
        for l <= mid && r <= right {
            if nums[l] < nums[r] {
                merged = append(merged, nums[l])
                l++
            } else {
                merged = append(merged, nums[r])
                r++
            }
        }
        for l <= mid {merged = append(merged, nums[l]);l++}
        for r <= right {merged = append(merged, nums[r]);r++}
        
        mPtr := 0
        i := left
        for mPtr < len(merged) {
            nums[i] = merged[mPtr]
            mPtr++
            i++
        }
    }
    dfs(0, len(nums)-1)
    return nums
}