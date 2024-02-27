func combinationSum(candidates []int, target int) [][]int {
    out := [][]int{}
    var dfs func(start int, t int, path []int)
    dfs = func(start, t int, path []int) {
        // base
        if t >= target {
            if t == target {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
            return
        }
        if start == len(candidates) {return}

        // logic
        // not choose
        dfs(start+1, t, path)
        // choose
        // action
        path = append(path, candidates[start])
        // recurse
        newSum := t+candidates[start]
        dfs(start, newSum, path)
        // backtrack
        path = path[:len(path)-1]
    }
    dfs(0,0,nil)
    return out
}