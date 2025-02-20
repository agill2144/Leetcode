func findDifferentBinaryString(nums []string) string {
    n := len(nums)
    set := map[string]bool{}
    visited := map[string]bool{}
    for i := 0; i < n; i++ {set[nums[i]] = true}
    choices := []string{"0","1"}
    var dfs func(path string) string
    dfs = func(path string) string {
        // base
        if len(path) == n {
            if !set[path] {return path}
            return ""
        }

        // logic
        for i := 0; i < len(choices); i++ {
            newP := path + choices[i]
            if !set[newP] && !visited[newP] {
                visited[newP] = true
                val := dfs(newP)
                if val != "" {return val}
                visited[newP] = false
            }
        }
        return ""
    }
    return dfs("")
}