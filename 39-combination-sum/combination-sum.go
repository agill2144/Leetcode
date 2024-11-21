func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(start int, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        // base
        if sum > target {return}

        // logic
        if sum == target {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }
        for i := start; i < len(candidates); i++ {
            if sum + candidates[i] <= target {
                sum += candidates[i]
                path = append(path, candidates[i])
                dfs(i, sum, path)
                path = path[:len(path)-1]
                sum -= candidates[i]
            }
        }
    }
    dfs(0, 0, nil)
    return out
}