func combinationSum3(k int, n int) [][]int {
    out := [][]int{}
    var dfs func(start, t int, path []int)
    dfs = func(start, t int, path []int){
        // base
        if t <= 0 {
            if t == 0 && len(path) == k {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        }
        if len(path) == k {return}
        
        // logic
        for i := start; i <= 9; i++ {
            path = append(path, i)
            dfs(i+1, t-i, path)
            path = path[:len(path)-1]
        }
    }
    dfs(1,n, []int{})
    return out
}