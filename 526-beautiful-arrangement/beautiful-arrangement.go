func countArrangement(n int) int {
    nums := []int{}
    for i := 0; i < n; i++ {nums = append(nums,i+1)}
    count := 0
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) {count++}

        // logic
        for i := start; i < len(nums); i++ {
            nums[i], nums[start] = nums[start], nums[i]
            if nums[start] % (start+1) == 0 || (start + 1) % nums[start] == 0 {
                dfs(start+1)
            }
            nums[i], nums[start] = nums[start], nums[i]
            
        }
    }
    dfs(0)
    return count
}