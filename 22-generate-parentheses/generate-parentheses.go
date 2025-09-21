func generateParenthesis(n int) []string {
	out := []string{}
	var dfs func(o, c int, path string, tmp []string)
	dfs = func(o, c int, path string, tmp []string) {
		// base
		if o == 0 && c == 0 {
			out = append(out, strings.Join(tmp, ""))
			return
		}
		// logic
		if o > 0 {
            // action
			tmp = append(tmp, "(")
            // recurse
			dfs(o-1, c, path+"(", tmp)
            // backtrack
			tmp = tmp[:len(tmp)-1]
		}

        // only close if we have opened more than we have closed
        // and thats only possible if closing avail count > opening avail count
        // meaning we have more close parans than open parans
        // i.e we have opened more than we have closed
        // i.e we can def close atleast one now ...
        // hence this logic
		if c > o {
			tmp = append(tmp, ")")
			dfs(o, c-1, path+")", tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(n, n, "", []string{})
	return out

}