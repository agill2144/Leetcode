func fib(n int) int {
    memo := map[int]int{}
    var dfs func(x int) int 
    dfs = func(x int) int {
        // base
        if x == 1 || x == 0  {return x}
        
        // logic
        if _, ok := memo[x-1]; !ok {
            memo[x-1] = dfs(x-1)
        }
        if _, ok := memo[x-2]; !ok {
            memo[x-2] = dfs(x-2)
        }
        return memo[x-1]+memo[x-2]
    }
    return dfs(n)
}