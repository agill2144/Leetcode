func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(start int, path []int, t int)
    dfs = func(start int, path []int, t int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out, newL)
            }
            return
        }
        if start == len(candidates) {return}
        
        
        // logic
        for i := start; i < len(candidates); i++ {
            path = append(path, candidates[i])
            dfs(i, path, t-candidates[i])
            path = path[:len(path)-1]
        }
    }
    dfs(0, nil, target)
    return out
}