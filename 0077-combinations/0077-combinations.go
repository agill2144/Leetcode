func combine(n int, k int) [][]int {
    out := [][]int{}
    var dfs func(start int, path []int)
    dfs = func(start int, path []int) {
        // base
        
        // logic
        if len(path) == k {
            newPath := make([]int, len(path))
            copy(newPath, path)
            out = append(out, newPath)
            return
        }
        
        for i := start; i <= n; i++ {
            path = append(path, i)
            dfs(i+1, path)
            path = path[:len(path)-1]
        }
        
    }
    dfs(1, []int{})
    return out
}