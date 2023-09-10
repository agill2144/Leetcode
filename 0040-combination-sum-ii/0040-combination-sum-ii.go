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
    var dfs func(start, t int, path []int)
    dfs = func(start, t int, path []int) {
        // base
        if t <= 0 {
            if t == 0 {
                newL := make([]int, len(path))
                copy(newL, path)
                out = append(out,newL)
            }
            return
        }
        
        // logic
        for i := start; i < len(deduped); i++ {
            val, count := deduped[i][0], deduped[i][1]
            if count == 0 {continue}
            
            // action
            path = append(path, val)
            deduped[i][1]--

            // recurse
            dfs(i, t-val, path)
            
            // backtrack
            // whether the path worked or not worked, undo it to use the next element instead
            // this is the sneaky "not-choose this number" case that we see in 0/1 recursion
            path = path[:len(path)-1]
            deduped[i][1]++            
        }
    }
    dfs(0, target, nil)
    return out
}