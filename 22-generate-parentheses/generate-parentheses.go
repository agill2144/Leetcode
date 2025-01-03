func generateParenthesis(n int) []string {
    out := []string{}
    var dfs func(o, c int, path string)
    dfs = func(o, c int, path string) {
        // base
        if o == 0 && c == 0 {
            out = append(out, path)
            return
        }

        // logic
        if o > 0 {
            dfs(o-1,c, path + "(")
        }
        // n = 3
        // ()() open = 2; close = 1
        //
        if c > o {
            dfs(o, c-1, path + ")")
        }
    }
    dfs(n,n,"")
    return out
}