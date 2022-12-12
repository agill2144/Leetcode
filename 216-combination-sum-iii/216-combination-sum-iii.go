func combinationSum3(k int, n int) [][]int {
    var result [][]int
    
    var dfs func(paths []int, target, start int)
    dfs = func(paths []int, target, start int) {
        
        // base
        // when this combination works
        if target == 0 && len(paths) == k {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        // when this combination does not work
        if target < 0 || start > 9 {
            return 
        }
        
        // logic
        for i := start; i <= 9; i++ {
            // action
            paths = append(paths, i)
            if len(paths) > k {return}
            // recurse to eval current combination
            dfs(paths, target-i, i+1) // i+1 because we cannot re-use the same number again.
            // backtrack : undo all the actions
            paths = paths[:len(paths)-1]
        }
        
    }
    
    dfs(nil, n, 1)
    
    return result
}