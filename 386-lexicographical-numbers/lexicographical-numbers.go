func lexicalOrder(n int) []int {
    out := []int{}
    var dfs func(curr int)
    dfs = func(curr int) {
        // base
        if curr > n {return}

        out = append(out, curr)
        // logic
        for i := 0; i <= 9; i++ {
            next := curr * 10 + i
            if next <= n {
                dfs(next)
            }
        }
    }
    for i := 1; i <= 9; i++ {
        dfs(i)
    }
    return out
}