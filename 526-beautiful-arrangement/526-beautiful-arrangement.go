func countArrangement(n int) int {
    
    nums := []int{}
    for i := 1; i <= n; i++ {
        nums = append(nums, i)
    }
    
    count := 0
    
    var backtrack func(start int)
    backtrack = func(start int) {
        // base
        if start == len(nums) {count++; return}
        
        // logic
        for i := start; i < len(nums); i++ {
            if (nums[i] % (start+1) == 0) || ((start+1) % nums[i] == 0) {
                // action
                nums[i], nums[start] = nums[start], nums[i]
                // recurse
                backtrack(start+1)
                // backtrack
                nums[i], nums[start] = nums[start], nums[i] 
            }
        }
    }
    
    backtrack(0)
    return count
}