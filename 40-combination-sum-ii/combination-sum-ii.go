func combinationSum2(candidates []int, target int) [][]int {
    freq := map[int]int{}
    for i := 0; i < len(candidates); i++ {
        freq[candidates[i]]++
    }
    deduped := [][]int{}
    for k, v := range freq {
        deduped = append(deduped, []int{k,v})
    }
    out := [][]int{}
    var dfs func(start int, sum int, path []int)
    dfs = func(start int, sum int, path []int) {
        // base
        if sum > target {return}

        // logic
        if sum == target {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }
        for i := start ; i < len(deduped); i++ {
            if deduped[i][1] > 0 {
                // action
                path = append(path, deduped[i][0])
                sum += deduped[i][0]
                deduped[i][1]--
                // recurse
                dfs(i, sum, path)
                // backtrack
                deduped[i][1]++
                sum -= deduped[i][0]
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0,0, nil)
    return out
}