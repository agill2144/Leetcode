func getHappyString(n int, k int) string {
    if n <= 1 && k > 3 {return ""}
    paths := []string{}
    choices := []string{"a","b","c"}
    var dfs func (path string) bool
    dfs = func (path string) bool{
        // base
        if len(path) >= n {
            if len(path) == n {
                paths = append(paths, path)
                if len(paths) == k {return true}
            }
            return false
        }

        // logic
        for i := 0; i < len(choices); i++ {
            curr := choices[i]
            if path == "" {
                if dfs(path + curr) {return true}
            } else {
                if curr != string(path[len(path)-1]) {
                    if dfs(path + curr) {return true}
                }
            }
        }
        return false
    }
    dfs("")
    if k-1 >= len(paths) {return ""}
    return paths[k-1]
}