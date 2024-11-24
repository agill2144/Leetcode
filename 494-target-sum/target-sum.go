func findTargetSumWays(nums []int, target int) int {
    count := 0
    var dfs func(i, sum int)
    dfs = func(i, sum int) {
        // base
        if i == len(nums) {
            if sum == target {count++}
            return
        }

        // logic
        dfs(i+1, sum+nums[i])
        dfs(i+1, sum-nums[i])
        
    }
    dfs(0, 0)
    return count
}