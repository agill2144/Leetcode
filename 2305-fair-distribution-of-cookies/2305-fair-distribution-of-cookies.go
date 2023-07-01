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
            oldMax := max
            if childs[i] > max {max = childs[i]}

            // recurse
            dfs(start, cookieIdx+1, childs, max)
            
            // backtrack
            childs[i] -= cookies[cookieIdx]
            max = oldMax
        }
    }
    dfs(0,0,make([]int, k),0)
    return out
}