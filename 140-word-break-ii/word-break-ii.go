func wordBreak(s string, wordDict []string) []string {
    set := map[string]struct{}{}
    for i := 0; i < len(wordDict); i++ {
        set[wordDict[i]] = struct{}{}
    }
    out := []string{}
    var dfs func(start int, path []string)
    dfs = func(start int, path []string) {
        // base
        if start == len(s) {
            pathStr := strings.Join(path, " ")
            out = append(out, pathStr)
            return
        }
         
        // logic
        for i := start ; i < len(s); i++ {
            subStr := s[start:i+1]
            if _, ok := set[subStr]; ok {
                path = append(path, subStr)
                dfs(i+1, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, nil)
    return out
}