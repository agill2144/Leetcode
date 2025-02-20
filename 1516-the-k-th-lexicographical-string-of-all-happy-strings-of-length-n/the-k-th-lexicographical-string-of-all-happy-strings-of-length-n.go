func getHappyString(n int, k int) string {
    if n <= 1 && k > 3 {return ""}
    paths := []string{}
    choices := []string{"a","b","c"}
    var dfs func (path string)
    dfs = func (path string){
        // base
        if len(path) >= n {
            if len(path) == n {
                paths = append(paths, path)
            }
            return
        }

        // logic
        for i := 0; i < len(choices); i++ {
            curr := choices[i]
            if path == "" {
                dfs(path + curr)
            } else {
                if curr != string(path[len(path)-1]) {
                    dfs(path + curr)
                }
            }
        }
    }
    dfs("")
    if k-1 >= len(paths) {return ""}
    return paths[k-1]
}