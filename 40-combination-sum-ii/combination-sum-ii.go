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
    var dfs func(start int, t int, path []int)
    dfs = func(start int, t int, path []int) {
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
        for i := start; i < len(deduped); i++ {
            if deduped[i][1] == 0 {continue}
            path = append(path, deduped[i][0])
            deduped[i][1]--
            dfs(i, t-deduped[i][0], path)
            deduped[i][1]++
            path = path[:len(path)-1]
        }
    }
    dfs(0, target, nil)
    return out
}