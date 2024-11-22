func combinationSum2(candidates []int, target int) [][]int {
    sort.Ints(candidates)
    out := [][]int{}

    var dfs func(start int, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        // base
        if sum > target {return}
        if sum == target {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }

        // logic
        for i := start; i < len(candidates); i++ {
            // skip if ith element == prev element ( to avoid duplicate combinations )
            // at every single start idx, we are starting a new path
            // when i == start, we want to use this numnber
            // when i > start then we dont want to use the curr number
            // if curr number is same as prev number
            // because if prev == curr, then at the prev idx, we have already used prev number
            // in the prev path, therefore dont use it again
            if i > start && candidates[i] == candidates[i-1] {continue}
            sum += candidates[i]
            path = append(path, candidates[i])
            dfs(i+1, sum, path)
            path = path[:len(path)-1]
            sum -= candidates[i]
        }
    }
    dfs(0,0, nil)
    return out
}