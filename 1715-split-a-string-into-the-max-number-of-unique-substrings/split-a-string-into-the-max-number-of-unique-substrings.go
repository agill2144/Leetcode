func maxUniqueSplit(s string) int {
    maxSplit := 0
    var dfs func(start int, path map[string]bool)
    dfs = func(start int, path map[string]bool) {
        // base
        if start == len(s) {
            maxSplit = max(maxSplit, len(path))
            return
        }

        // short circuit 
        if len(s)-start + len(path) <= maxSplit {return}

        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if !path[subStr] {
                path[subStr] = true
                dfs(i+1, path)
                delete(path, subStr)
            }
        }
    }
    dfs(0, map[string]bool{})
    return maxSplit
}