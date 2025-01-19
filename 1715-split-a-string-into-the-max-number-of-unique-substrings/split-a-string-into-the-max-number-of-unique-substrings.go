func maxUniqueSplit(s string) int {
    maxCount := 0
    var dfs func(start int, path map[string]bool)
    dfs = func(start int, path map[string]bool) {
        // base
        if start == len(s) {
            maxCount = max(maxCount, len(path))
            return
        }
        if len(path) + len(s)-start <= maxCount {return}
        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if path[subStr] {continue}
            path[subStr] = true
            dfs(i+1, path)
            delete(path, subStr)
        }
    }
    dfs(0, map[string]bool{})
    return maxCount
}