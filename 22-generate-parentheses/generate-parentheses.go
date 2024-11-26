func generateParenthesis(n int) []string {
    out := []string{}
    var dfs func(o, c int, path string)
    dfs = func(o, c int, path string) {
        // base
        if o == 0 && c == 0 {out = append(out, path); return}

        // logic
        if o > 0 {
            dfs(o-1, c, path + "(")
        }
        // there must be open parans for us to close
        // and thats only possible if open count was less than close count
        // meaning we have more close parans left than open
        // meaning there are open parans that have yet to be closed, therefore start closing them
        if c > o {
            dfs(o, c-1, path + ")")
        }
    } 
    dfs(n, n, "")
    return out
}