func countArrangement(n int) int {
    nums := make([]int, n+1)
    for i := 1; i < len(nums); i++ {
        nums[i] = i
    }
    count := 0
    
    var dfs func(start int)
    dfs = func(start int) {
        // base
        if start == len(nums) { count++; return}
        
        // logic
        for i := start; i < len(nums); i++ {
            nums[i], nums[start] = nums[start] , nums[i]
            if (nums[start] % start == 0 || start % nums[start] == 0) {
                dfs(start+1)
            }
            nums[i], nums[start] = nums[start] , nums[i]
        }
    }
    
    dfs(1)
    return count
}