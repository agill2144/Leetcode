func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(t , start int, path []int)
    dfs = func(t, start int, path []int) {
        // base
        if t == 0 {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }
        if t < 0 || start == len(candidates) {return}
        
        // logic
        for i := start; i < len(candidates); i++ {
            path = append(path, candidates[i])
            dfs(t-candidates[i], i, path)
            path = path[:len(path)-1]
        }
    }
    dfs(target, 0, nil)
    return out
}