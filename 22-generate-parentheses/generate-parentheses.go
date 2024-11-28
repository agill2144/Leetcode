func generateParenthesis(n int) []string {
    out := []string{}
    var dfs func(o, c int, path string)
    dfs = func(o, c int, path string) {
        // base
        if o == n && c == n {
            out = append(out, path)
            return
        }

        // logic
        if o < n {
            dfs(o+1, c, path + "(")
        }
        if c < o {
            dfs(o, c+1, path + ")")
        }
    }
    dfs(0,0,"")
    return out
}