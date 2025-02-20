func getHappyString(n int, k int) string {
    if n <= 1 && k > 3 {return ""}
    count := 0
    var dfs func (path string) string
    dfs = func (path string) string{
        // base
        if len(path) >= n {
            if len(path) == n {
                count++
                if count == k {return path}
            }
            return ""
        }

        // logic
        for i := 0; i < 3; i++ {
            curr := string('a' + i)
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