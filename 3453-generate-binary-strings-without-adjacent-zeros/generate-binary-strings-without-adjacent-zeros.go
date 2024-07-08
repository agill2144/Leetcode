func validStrings(n int) []string {
    choices := []int{'1','0'}
    out := []string{}
    var dfs func(path string)
    dfs = func(path string) {
        // base
        if len(path) == n {
            out = append(out, path)
            return
        }

        // logic
        for i := 0; i < len(choices); i++ {
            char := choices[i]
            path += string(char)
            if len(path) >= 2 {
                if path[len(path)-1] == '1' || path[len(path)-2] == '1' {
                    dfs(path)
                }
            } else {
                dfs(path)
            }
            path = path[:len(path)-1]
        }
    }
    dfs("")
    return out
}