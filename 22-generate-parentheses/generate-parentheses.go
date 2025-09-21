func generateParenthesis(n int) []string {
	out := []string{}
	var dfs func(o, c int, path string )
	dfs = func(o, c int, path string) {
		// base
		if o == 0 && c == 0 {
			out = append(out, path)
			return
		}
		// logic
		if o > 0 {
            // action
            path += "("
            // recurse
			dfs(o-1, c, path)
            // backtrack
            path = path[:len(path)-1]
		}

        // only close if we have opened more than we have closed
        // and thats only possible if closing avail count > opening avail count
        // meaning we have more close parans than open parans
        // i.e we have opened more than we have closed
        // i.e we can def close atleast one now ...
        // hence this logic
		if c > o {
            path += ")"
            // recurse
			dfs(o, c-1, path)
            // backtrack
            path = path[:len(path)-1]
		}
	}
	dfs(n, n, "")
	return out

}