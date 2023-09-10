func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(start, t int, path []int)
    dfs = func(start, t int, path []int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(candidates); i++ {
            path = append(path, candidates[i])
            dfs(i, t-candidates[i], path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, target, []int{})
    return out
}