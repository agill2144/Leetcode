func restoreIpAddresses(s string) []string {
    out := []string{}
    var dfs func(start int, path []string)
    dfs = func(start int, path []string) {
        // base
        if start == len(s) {
            if len(path) == 4 {
                out = append(out, strings.Join(path, "."))
            }
            return
        }

        // logic
        for i := start; i < len(s); i++ {
            subStr := s[start:i+1]
            if len(subStr) > 1 && subStr[0] == '0' {return}
            subStrInt, _ := strconv.Atoi(subStr)
            if subStrInt >= 0 && subStrInt <= 255 {
                path = append(path, subStr)
                dfs(i+1, path)
                path = path[:len(path)-1]
            }
        }
    }
    dfs(0, nil)
    return out
}