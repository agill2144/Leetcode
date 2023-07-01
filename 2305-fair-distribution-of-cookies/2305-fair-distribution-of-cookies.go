func distributeCookies(cookies []int, k int) int {
    out := math.MaxInt64
    var dfs func(start int, cookieIdx int, childs []int, max int)
    dfs = func(start int, cookieIdx int, childs []int, max int) {
        // base
        if cookieIdx == len(cookies) {
            if max < out {out = max}
            return
        }
        
        // logic
        for i := start; i < len(childs); i++ {
            // action
            childs[i] += cookies[cookieIdx]
            
            // we could just update the max arg, but then we would have to save the old value and backtrack
            newMax := max
            if childs[i] > newMax {newMax = childs[i]}

            // recurse
            dfs(start, cookieIdx+1, childs, newMax)
            
            // backtrack
            childs[i] -= cookies[cookieIdx]
        }
    }
    dfs(0,0,make([]int, k),0)
    return out
}