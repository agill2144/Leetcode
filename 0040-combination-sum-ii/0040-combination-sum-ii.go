// n = len(candidates)
// t = target
func combinationSum2(candidates []int, target int) [][]int {

    // space = o(n)
    freq := map[int]int{}
        
    // time = o(n)
    for i := 0; i < len(candidates); i++ {
        freq[candidates[i]]++
    }
    
    // space = o(n)
    deduped := [][]int{}

    // time = o(n)
    for k, v := range freq {
        deduped = append(deduped, []int{k,v})
    }
    
    out := [][]int{}
    var dfs func(start int, t int, path []int)

    
    // time = 2^(t+n)
    // space = t for recursive stack * t for new path ; t^2
    dfs = func(start, t int, path []int) {
        // base
        if t <= 0 {
            if t == 0 {
                newP := make([]int, len(path))
                copy(newP, path)
                out = append(out, newP)
            }
            return
        }
        
        // logic
        for i := start; i < len(deduped); i++ {
            num, count := deduped[i][0], deduped[i][1]
            if count == 0 {continue}
            // action
            path = append(path, num)
            deduped[i][1]--
            
            // recurse
            dfs(i, t-num,path)

            // backtrack
            path = path[:len(path)-1]
            deduped[i][1]++
        }
    }
    dfs(0, target, nil)
    return out
}