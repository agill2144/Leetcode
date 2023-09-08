func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(ptr int, t int, path []int)
    dfs = func(ptr int, t int, path []int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        }
        if ptr == len(candidates) || t < 0 {return}
        
        // logic
        dfs(ptr+1, t, path)
        path = append(path, candidates[ptr])
        dfs(ptr, t-candidates[ptr], path)
        path = path[:len(path)-1]
    }
    dfs(0, target, []int{})
    return out
}