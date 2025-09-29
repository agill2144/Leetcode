func maxUniqueSplit(s string) int {
    if s == "" {return 0}
    ans := 0
    var dfs func(start int, paths map[string]bool)
    dfs = func(start int, paths map[string]bool) {
        // base
        if start == len(s) {
            ans = max(ans, len(paths))
            return
        }

        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if !paths[subStr] {
                paths[subStr] = true
                dfs(i+1, paths)
                delete(paths, subStr)
            }
        }
    }
    dfs(0, map[string]bool{})
    return ans
}