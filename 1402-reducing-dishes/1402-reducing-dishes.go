func maxSatisfaction(satisfaction []int) int {
    sort.Ints(satisfaction)
    m := len(satisfaction)+1
    n := len(satisfaction)+1
    memo := make([][]int, m)
    for i := 0; i < m; i++ {
        memo[i] = make([]int, n)
        for j := 0; j < n; j++ {
            memo[i][j] = -1
        }
    }
    var dfs func(idx, time int) int    
    dfs = func(idx, time int) int {
        // base
        if idx == len(satisfaction) {
            return 0
        }
        
        if memo[idx][time] != -1 {return memo[idx][time]}
        
        
        // logic
        i := (satisfaction[idx]*time) + dfs(idx+1, time+1)
        e := dfs(idx+1, time)
        memo[idx][time] = max(i,e)
        
        return max(i,e) 
    }
    
   
    return  dfs(0,1)
}

func max(x, y int) int {
    if x > y {return x}
    return y
}