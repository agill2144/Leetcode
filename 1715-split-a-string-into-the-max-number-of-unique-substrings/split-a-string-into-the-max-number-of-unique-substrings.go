func maxUniqueSplit(s string) int {
    n := len(s)
    set := map[string]bool{}
    maxSplits := 0
    var dfs func(start int)
    // ababccc
    dfs = func(start int) {
        // base
        if start == n {
            maxSplits = max(maxSplits, len(set))
            return
        }

        // logic
        for i := start; i < n; i++ {
            subStr := s[start:i+1]
            if !set[subStr] {
                // action
                set[subStr] = true
                // recurse
                dfs(i+1)
                // backtrack
                delete(set, subStr)
            }
        }
    }
    dfs(0)
    return maxSplits
}