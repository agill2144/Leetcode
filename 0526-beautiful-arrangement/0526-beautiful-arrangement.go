func countArrangement(n int) int {
    
    nums := make([]int, n+1)
    for i := 1; i < len(nums); i++ {
        nums[i] = i
    }
    
    count := 0
    
    var backtrack func(start int)
    backtrack = func(start int) {
        // base
        if start == len(nums) {count++; return}
        
        // logic
        for i := start; i < len(nums); i++ {
            if (nums[i] % (start) == 0) || ((start) % nums[i] == 0) {
                // action
                nums[i], nums[start] = nums[start], nums[i]
                // recurse
                backtrack(start+1)
                // backtrack
                nums[i], nums[start] = nums[start], nums[i] 
            }
        }
    }
    
    backtrack(1)
    return count
}