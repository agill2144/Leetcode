// time: o(n) + o(n!)
// space: o(n) nums arr + o(n) recursion stack
func countArrangement(n int) int {
    
    // o(n) time
    // o(n) space
    nums := make([]int, n+1)
    for i := 1; i < len(nums); i++ {
        nums[i] = i
    }
    count := 0
    
    var dfs func(start int)
    // permutations take worst case o(n!) time
    dfs = func(start int) {
        // base
        if start == len(nums) { count++; return}
        
        // logic
        for i := start; i < len(nums); i++ {
            // permute / action 
            nums[i], nums[start] = nums[start] , nums[i]
            // check / validate
            if (nums[start] % start == 0 || start % nums[start] == 0) {
                // recurse
                dfs(start+1)
            }
            // backtrack
            nums[i], nums[start] = nums[start] , nums[i]
        }
    }
    
    dfs(1)
    return count
}