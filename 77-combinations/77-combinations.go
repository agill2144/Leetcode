func combine(n int, k int) [][]int {
    var result [][]int
    
    var dfs func(paths []int, start int) 
    dfs = func(paths []int, start int) {
        // base
        if len(paths) == k {
            newL := make([]int, len(paths))
            copy(newL, paths)
            result = append(result, newL)
            return
        }
        if start > n {
            return
        }
        
        // logic
        for i := start;i <= n; i++{
            // action
            paths = append(paths, i)
            // recurse
            dfs(paths, i+1)
            // backtrack
            paths = paths[:len(paths)-1]
        }
    }
    dfs(nil, 1)
    return result
}