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
            if len(path) < 2 {
                // too small to validate subtring
                dfs(path)
            } else {
                // only proceed with this path, IF we have 1 in either of the last 2 positions
                // we do not need to validate the entire substr/path each time,
                // as long as we have a 1 in either last or 2nd last position, we can continue with this path
                if path[len(path)-1] == '1' || path[len(path)-2] == '1' {
                    dfs(path)
                }
            }
            path = path[:len(path)-1]
        }
    }
    dfs("")
    return out
}