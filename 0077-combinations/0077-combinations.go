func combine(n int, k int) [][]int {
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base        
        if len(path) == k {
            newPath := make([]int, k)
            copy(newPath, path)
            out = append(out, newPath)
            return
        }
        if start > n {return}
        
        // logic        
        for i := start; i <= n ; i++ {
            // action
            path = append(path, i)
            // recurse
            dfs(i+1, path)
            // backtrack
            path = path[:len(path)-1]
        }
    }
    dfs(1, []int{})
    return out
}