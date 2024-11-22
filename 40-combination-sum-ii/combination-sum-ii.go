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
            /*
                why arent we comparing ith element with prev ?
                if curr and prev are same, we can skip curr
                - the reason is, prev element when recursion comes back
                is removed/backtracked!
                - the only element that does not get removed is the start element
                at this recursion node
                - therefore compare ith element with start element
                what if start == i ? then they are def the same
                - therefore add another check that i must be > start position
                - only then compare ith element with start element
            */
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