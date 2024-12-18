// numberOfOptionsPerNode^totalNumberOfOptions * extraWorkPerRecursion
// numberOfOptionsPerNode = 2
// totalNumberOfOptions = n
// extraWorkPerRecursion = 0
// tc = o(2^n)
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
            dfs(o-1, c, path + "(")
        }
        if c > 0 && o < c {
            dfs(o, c-1, path + ")")
        }
    }
    dfs(n, n, "")
    return out
}