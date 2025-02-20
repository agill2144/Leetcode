func getHappyString(n int, k int) string {
    if n <= 1 && k > 3 {return ""}
    paths := []string{}
    count := 0
    choices := []string{"a","b","c"}
    var dfs func (path string) string
    dfs = func (path string) string{
        // base
        if len(path) >= n {
            if len(path) == n {
                count++
                paths = append(paths, path)
                // if len(paths) == k {return path}
                if count == k {return path}
            }
            return ""
        }

        // logic
        for i := 0; i < len(choices); i++ {
            curr := choices[i]
            prev := ""
            if path != "" {prev = string(path[len(path)-1])}
            if curr != prev {
                val := dfs(path + curr) 
                if val != "" {return val}
            }
        }
        return ""
    }
    return dfs("")
}