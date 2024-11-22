func combinationSum3(k int, n int) [][]int {
    out := [][]int{}
    var dfs func(start, sum int, path []int)
    dfs = func(start, sum int, path []int) {
        // base
        if sum > n || len(path) > k {return}

        // logic
        if sum == n && len(path) == k {
            newL := make([]int, len(path))
            copy(newL, path)
            out = append(out, newL)
            return
        }
        for i := start; i <= 9; i++ {
            sum += i
            path = append(path, i)
            dfs(i+1, sum, path)
            path = path[:len(path)-1]
            sum -= i
        } 
    }
    dfs(1, 0, nil)
    return out
}