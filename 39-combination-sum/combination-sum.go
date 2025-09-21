func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(start int, sum int, path []int)
    dfs = func(start int, sum int, path []int) {
        // base
        if sum >= target {
            if sum == target {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
            return
        }

        // logic
        for i := start; i < len(candidates); i++ {
            // action
            sum += candidates[i]
            path = append(path, candidates[i])
            // recurse
            dfs(i, sum, path)
            // backtrack
            path = path[:len(path)-1]
            sum -= candidates[i]
        }
    }
    dfs(0,0,nil)
    return out
}